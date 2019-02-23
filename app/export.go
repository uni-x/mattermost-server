// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/uni-x/mattermost-server/model"
	"github.com/pkg/errors"
)

func (a *App) BulkExport(writer io.Writer, file string, pathToEmojiDir string, dirNameToExportEmoji string) *model.AppError {
	if err := a.ExportVersion(writer); err != nil {
		return err
	}

	if err := a.ExportAllTeams(writer); err != nil {
		return err
	}

	if err := a.ExportAllChannels(writer); err != nil {
		return err
	}

	if err := a.ExportAllUsers(writer); err != nil {
		return err
	}

	if err := a.ExportAllPosts(writer); err != nil {
		return err
	}
	if err := a.ExportCustomEmoji(writer, file, pathToEmojiDir, dirNameToExportEmoji); err != nil {
		return err
	}

	return nil
}

func (a *App) ExportWriteLine(writer io.Writer, line *LineImportData) *model.AppError {
	b, err := json.Marshal(line)
	if err != nil {
		return model.NewAppError("BulkExport", "app.export.export_write_line.json_marshall.error", nil, "err="+err.Error(), http.StatusBadRequest)
	}

	if _, err := writer.Write(append(b, '\n')); err != nil {
		return model.NewAppError("BulkExport", "app.export.export_write_line.io_writer.error", nil, "err="+err.Error(), http.StatusBadRequest)
	}

	return nil
}

func (a *App) ExportVersion(writer io.Writer) *model.AppError {
	version := 1
	versionLine := &LineImportData{
		Type:    "version",
		Version: &version,
	}

	return a.ExportWriteLine(writer, versionLine)
}

func (a *App) ExportAllTeams(writer io.Writer) *model.AppError {
	afterId := strings.Repeat("0", 26)
	for {
		result := <-a.Srv.Store.Team().GetAllForExportAfter(1000, afterId)

		if result.Err != nil {
			return result.Err
		}

		teams := result.Data.([]*model.TeamForExport)

		if len(teams) == 0 {
			break
		}

		for _, team := range teams {
			afterId = team.Id

			// Skip deleted.
			if team.DeleteAt != 0 {
				continue
			}

			teamLine := ImportLineFromTeam(team)
			if err := a.ExportWriteLine(writer, teamLine); err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *App) ExportAllChannels(writer io.Writer) *model.AppError {
	afterId := strings.Repeat("0", 26)
	for {
		result := <-a.Srv.Store.Channel().GetAllChannelsForExportAfter(1000, afterId)

		if result.Err != nil {
			return result.Err
		}

		channels := result.Data.([]*model.ChannelForExport)

		if len(channels) == 0 {
			break
		}

		for _, channel := range channels {
			afterId = channel.Id

			// Skip deleted.
			if channel.DeleteAt != 0 {
				continue
			}

			channelLine := ImportLineFromChannel(channel)
			if err := a.ExportWriteLine(writer, channelLine); err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *App) ExportAllUsers(writer io.Writer) *model.AppError {
	afterId := strings.Repeat("0", 26)
	for {
		result := <-a.Srv.Store.User().GetAllAfter(1000, afterId)

		if result.Err != nil {
			return result.Err
		}

		users := result.Data.([]*model.User)

		if len(users) == 0 {
			break
		}

		for _, user := range users {
			afterId = user.Id

			// Skip deleted.
			if user.DeleteAt != 0 {
				continue
			}

			userLine := ImportLineFromUser(user)

			userLine.User.NotifyProps = a.buildUserNotifyProps(user.NotifyProps)

			// Do the Team Memberships.
			members, err := a.buildUserTeamAndChannelMemberships(user.Id)
			if err != nil {
				return err
			}

			userLine.User.Teams = members

			if err := a.ExportWriteLine(writer, userLine); err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *App) buildUserTeamAndChannelMemberships(userId string) (*[]UserTeamImportData, *model.AppError) {
	var memberships []UserTeamImportData

	result := <-a.Srv.Store.Team().GetTeamMembersForExport(userId)

	if result.Err != nil {
		return nil, result.Err
	}

	members := result.Data.([]*model.TeamMemberForExport)

	for _, member := range members {
		// Skip deleted.
		if member.DeleteAt != 0 {
			continue
		}

		memberData := ImportUserTeamDataFromTeamMember(member)

		// Do the Channel Memberships.
		channelMembers, err := a.buildUserChannelMemberships(userId, member.TeamId)
		if err != nil {
			return nil, err
		}

		memberData.Channels = channelMembers

		memberships = append(memberships, *memberData)
	}

	return &memberships, nil
}

func (a *App) buildUserChannelMemberships(userId string, teamId string) (*[]UserChannelImportData, *model.AppError) {
	var memberships []UserChannelImportData

	result := <-a.Srv.Store.Channel().GetChannelMembersForExport(userId, teamId)
	if result.Err != nil {
		return nil, result.Err
	}

	members := result.Data.([]*model.ChannelMemberForExport)

	category := model.PREFERENCE_CATEGORY_FAVORITE_CHANNEL
	preferences, err := a.GetPreferenceByCategoryForUser(userId, category)
	if err != nil && err.StatusCode != http.StatusNotFound {
		return nil, err
	}

	for _, member := range members {
		memberships = append(memberships, *ImportUserChannelDataFromChannelMemberAndPreferences(member, &preferences))
	}
	return &memberships, nil
}

func (a *App) buildUserNotifyProps(notifyProps model.StringMap) *UserNotifyPropsImportData {

	getProp := func(key string) *string {
		if v, ok := notifyProps[key]; ok {
			return &v
		}
		return nil
	}

	return &UserNotifyPropsImportData{
		Desktop:          getProp(model.DESKTOP_NOTIFY_PROP),
		DesktopSound:     getProp(model.DESKTOP_SOUND_NOTIFY_PROP),
		Email:            getProp(model.EMAIL_NOTIFY_PROP),
		Mobile:           getProp(model.PUSH_NOTIFY_PROP),
		MobilePushStatus: getProp(model.PUSH_STATUS_NOTIFY_PROP),
		ChannelTrigger:   getProp(model.CHANNEL_MENTIONS_NOTIFY_PROP),
		CommentsTrigger:  getProp(model.COMMENTS_NOTIFY_PROP),
		MentionKeys:      getProp(model.MENTION_KEYS_NOTIFY_PROP),
	}
}

func (a *App) ExportAllPosts(writer io.Writer) *model.AppError {
	afterId := strings.Repeat("0", 26)
	for {
		result := <-a.Srv.Store.Post().GetParentsForExportAfter(1000, afterId)

		if result.Err != nil {
			return result.Err
		}

		posts := result.Data.([]*model.PostForExport)

		if len(posts) == 0 {
			break
		}

		for _, post := range posts {
			afterId = post.Id

			// Skip deleted.
			if post.DeleteAt != 0 {
				continue
			}

			postLine := ImportLineForPost(post)

			// Do the Replies.
			replies, err := a.buildPostReplies(post.Id)
			if err != nil {
				return err
			}

			reactions, err := a.BuildPostReactions(post.Id)
			if err != nil {
				return err
			}

			postLine.Post.Replies = replies

			postLine.Post.Reactions = reactions

			if err := a.ExportWriteLine(writer, postLine); err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *App) buildPostReplies(postId string) (*[]ReplyImportData, *model.AppError) {
	var replies []ReplyImportData

	result := <-a.Srv.Store.Post().GetRepliesForExport(postId)

	if result.Err != nil {
		return nil, result.Err
	}

	replyPosts := result.Data.([]*model.ReplyForExport)

	for _, reply := range replyPosts {
		replyImportObject := ImportReplyFromPost(reply)
		if reply.HasReactions == true {
			reactionsOfReply, err := a.BuildPostReactions(reply.Id)
			if err != nil {
				return nil, err
			}
			replyImportObject.Reactions = reactionsOfReply
		}
		replies = append(replies, *replyImportObject)
	}

	return &replies, nil
}

func (a *App) BuildPostReactions(postId string) (*[]ReactionImportData, *model.AppError) {
	var reactionsOfPost []ReactionImportData

	result := <-a.Srv.Store.Reaction().GetForPost(postId, true)
	if result.Err != nil {
		return nil, result.Err
	}

	reactions := result.Data.([]*model.Reaction)

	for _, reaction := range reactions {
		result := <-a.Srv.Store.User().Get(reaction.UserId)
		if result.Err != nil {
			return nil, result.Err
		}
		user := result.Data.(*model.User)
		reactionsOfPost = append(reactionsOfPost, *ImportReactionFromPost(user, reaction))
	}

	return &reactionsOfPost, nil

}

func (a *App) ExportCustomEmoji(writer io.Writer, file string, pathToEmojiDir string, dirNameToExportEmoji string) *model.AppError {
	pageNumber := 0
	for {
		customEmojiList, err := a.GetEmojiList(pageNumber, 100, model.EMOJI_SORT_BY_NAME)

		if err != nil {
			return err
		}

		if len(customEmojiList) == 0 {
			break
		}

		pageNumber++

		pathToDir := a.createDirForEmoji(file, dirNameToExportEmoji)

		for _, emoji := range customEmojiList {
			emojiImagePath := pathToEmojiDir + emoji.Id + "/image"
			err := a.copyEmojiImages(emoji.Id, emojiImagePath, pathToDir)
			if err != nil {
				return model.NewAppError("BulkExport", "app.export.export_custom_emoji.copy_emoji_images.error", nil, "err="+err.Error(), http.StatusBadRequest)
			}

			filePath := dirNameToExportEmoji + "/" + emoji.Id + "/image"

			emojiImportObject := ImportLineFromEmoji(emoji, filePath)

			if err := a.ExportWriteLine(writer, emojiImportObject); err != nil {
				return err
			}
		}
	}

	return nil
}

// Creates directory named 'exported_emoji' to copy the emoji files
// Directory and the file specified by admin share the same path
func (a *App) createDirForEmoji(file string, dirName string) string {
	pathToFile, _ := filepath.Abs(file)
	pathSlice := strings.Split(pathToFile, "/")
	if len(pathSlice) > 0 {
		pathSlice = pathSlice[:len(pathSlice)-1]
	}
	pathToDir := strings.Join(pathSlice, "/") + "/" + dirName

	if _, err := os.Stat(pathToDir); os.IsNotExist(err) {
		os.Mkdir(pathToDir, os.ModePerm)
	}
	return pathToDir
}

// Copies emoji files from 'data/emoji' dir to 'exported_emoji' dir
func (a *App) copyEmojiImages(emojiId string, emojiImagePath string, pathToDir string) error {
	fromPath, err := os.Open(emojiImagePath)
	if fromPath == nil || err != nil {
		return errors.New("Error reading " + emojiImagePath + "file")
	}
	defer fromPath.Close()

	emojiDir := pathToDir + "/" + emojiId

	if _, err = os.Stat(emojiDir); err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrapf(err, "Error fetching file info of emoji directory %v", emojiDir)
		}

		if err = os.Mkdir(emojiDir, os.ModePerm); err != nil {
			return errors.Wrapf(err, "Error creating emoji directory %v", emojiDir)
		}
	}

	toPath, err := os.OpenFile(emojiDir+"/image", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return errors.New("Error creating the image file " + err.Error())
	}
	defer toPath.Close()

	_, err = io.Copy(toPath, fromPath)
	if err != nil {
		return errors.New("Error copying emojis " + err.Error())
	}

	return nil
}
