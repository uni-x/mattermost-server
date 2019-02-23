// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uni-x/mattermost-server/model"
	"github.com/uni-x/mattermost-server/store"
	"github.com/uni-x/mattermost-server/store/storetest"
)

func TestCreatePostDeduplicate(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	t.Run("duplicate create post is idempotent", func(t *testing.T) {
		pendingPostId := model.NewId()
		post, err := th.App.CreatePostAsUser(&model.Post{
			UserId:        th.BasicUser.Id,
			ChannelId:     th.BasicChannel.Id,
			Message:       "message",
			PendingPostId: pendingPostId,
		}, false)
		require.Nil(t, err)
		require.Equal(t, "message", post.Message)

		duplicatePost, err := th.App.CreatePostAsUser(&model.Post{
			UserId:        th.BasicUser.Id,
			ChannelId:     th.BasicChannel.Id,
			Message:       "message",
			PendingPostId: pendingPostId,
		}, false)
		require.Nil(t, err)
		require.Equal(t, post.Id, duplicatePost.Id, "should have returned previously created post id")
		require.Equal(t, "message", duplicatePost.Message)
	})

	t.Run("post rejected by plugin leaves cache ready for non-deduplicated try", func(t *testing.T) {
		setupPluginApiTest(t, `
			package main

			import (
				"github.com/uni-x/mattermost-server/plugin"
				"github.com/uni-x/mattermost-server/model"
			)

			type MyPlugin struct {
				plugin.MattermostPlugin
				allow bool
			}

			func (p *MyPlugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
				if !p.allow {
					p.allow = true
					return nil, "rejected"
				}

				return nil, ""
			}

			func main() {
				plugin.ClientMain(&MyPlugin{})
			}
		`, `{"id": "testrejectfirstpost", "backend": {"executable": "backend.exe"}}`, "testrejectfirstpost", th.App)

		pendingPostId := model.NewId()
		post, err := th.App.CreatePostAsUser(&model.Post{
			UserId:        th.BasicUser.Id,
			ChannelId:     th.BasicChannel.Id,
			Message:       "message",
			PendingPostId: pendingPostId,
		}, false)
		require.NotNil(t, err)
		require.Equal(t, "Post rejected by plugin. rejected", err.Id)
		require.Nil(t, post)

		duplicatePost, err := th.App.CreatePostAsUser(&model.Post{
			UserId:        th.BasicUser.Id,
			ChannelId:     th.BasicChannel.Id,
			Message:       "message",
			PendingPostId: pendingPostId,
		}, false)
		require.Nil(t, err)
		require.Equal(t, "message", duplicatePost.Message)
	})

	t.Run("slow posting after cache entry blocks duplicate request", func(t *testing.T) {
		setupPluginApiTest(t, `
			package main

			import (
				"github.com/uni-x/mattermost-server/plugin"
				"github.com/uni-x/mattermost-server/model"
				"time"
			)

			type MyPlugin struct {
				plugin.MattermostPlugin
				instant bool
			}

			func (p *MyPlugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
				if !p.instant {
					p.instant = true
					time.Sleep(3 * time.Second)
				}

				return nil, ""
			}

			func main() {
				plugin.ClientMain(&MyPlugin{})
			}
		`, `{"id": "testdelayfirstpost", "backend": {"executable": "backend.exe"}}`, "testdelayfirstpost", th.App)

		var post *model.Post
		pendingPostId := model.NewId()

		wg := sync.WaitGroup{}

		// Launch a goroutine to make the first CreatePost call that will get delayed
		// by the plugin above.
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			post, err = th.App.CreatePostAsUser(&model.Post{
				UserId:        th.BasicUser.Id,
				ChannelId:     th.BasicChannel.Id,
				Message:       "plugin delayed",
				PendingPostId: pendingPostId,
			}, false)
			require.Nil(t, err)
			require.Equal(t, post.Message, "plugin delayed")
		}()

		// Give the goroutine above a chance to start and get delayed by the plugin.
		time.Sleep(2 * time.Second)

		// Try creating a duplicate post
		duplicatePost, err := th.App.CreatePostAsUser(&model.Post{
			UserId:        th.BasicUser.Id,
			ChannelId:     th.BasicChannel.Id,
			Message:       "plugin delayed",
			PendingPostId: pendingPostId,
		}, false)
		require.NotNil(t, err)
		require.Equal(t, "api.post.deduplicate_create_post.pending", err.Id)
		require.Nil(t, duplicatePost)

		// Wait for the first CreatePost to finish to ensure assertions are made.
		wg.Wait()
	})

	t.Run("duplicate create post after cache expires is not idempotent", func(t *testing.T) {
		pendingPostId := model.NewId()
		post, err := th.App.CreatePostAsUser(&model.Post{
			UserId:        th.BasicUser.Id,
			ChannelId:     th.BasicChannel.Id,
			Message:       "message",
			PendingPostId: pendingPostId,
		}, false)
		require.Nil(t, err)
		require.Equal(t, "message", post.Message)

		time.Sleep(PENDING_POST_IDS_CACHE_TTL)

		duplicatePost, err := th.App.CreatePostAsUser(&model.Post{
			UserId:        th.BasicUser.Id,
			ChannelId:     th.BasicChannel.Id,
			Message:       "message",
			PendingPostId: pendingPostId,
		}, false)
		require.Nil(t, err)
		require.NotEqual(t, post.Id, duplicatePost.Id, "should have created new post id")
		require.Equal(t, "message", duplicatePost.Message)
	})
}

func TestAttachFilesToPost(t *testing.T) {
	t.Run("should attach files", func(t *testing.T) {
		th := Setup().InitBasic()
		defer th.TearDown()

		info1 := store.Must(th.App.Srv.Store.FileInfo().Save(&model.FileInfo{
			CreatorId: th.BasicUser.Id,
			Path:      "path.txt",
		})).(*model.FileInfo)
		info2 := store.Must(th.App.Srv.Store.FileInfo().Save(&model.FileInfo{
			CreatorId: th.BasicUser.Id,
			Path:      "path.txt",
		})).(*model.FileInfo)

		post := th.BasicPost
		post.FileIds = []string{info1.Id, info2.Id}

		err := th.App.attachFilesToPost(post)
		assert.Nil(t, err)

		infos, err := th.App.GetFileInfosForPost(post.Id)
		assert.Nil(t, err)
		assert.Len(t, infos, 2)
	})

	t.Run("should update File.PostIds after failing to add files", func(t *testing.T) {
		th := Setup().InitBasic()
		defer th.TearDown()

		info1 := store.Must(th.App.Srv.Store.FileInfo().Save(&model.FileInfo{
			CreatorId: th.BasicUser.Id,
			Path:      "path.txt",
			PostId:    model.NewId(),
		})).(*model.FileInfo)
		info2 := store.Must(th.App.Srv.Store.FileInfo().Save(&model.FileInfo{
			CreatorId: th.BasicUser.Id,
			Path:      "path.txt",
		})).(*model.FileInfo)

		post := th.BasicPost
		post.FileIds = []string{info1.Id, info2.Id}

		err := th.App.attachFilesToPost(post)
		assert.Nil(t, err)

		infos, err := th.App.GetFileInfosForPost(post.Id)
		assert.Nil(t, err)
		assert.Len(t, infos, 1)
		assert.Equal(t, info2.Id, infos[0].Id)

		updated, err := th.App.GetSinglePost(post.Id)
		require.Nil(t, err)
		assert.Len(t, updated.FileIds, 1)
		assert.Contains(t, updated.FileIds, info2.Id)
	})
}

func TestUpdatePostEditAt(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	post := &model.Post{}
	*post = *th.BasicPost

	post.IsPinned = true
	if saved, err := th.App.UpdatePost(post, true); err != nil {
		t.Fatal(err)
	} else if saved.EditAt != post.EditAt {
		t.Fatal("shouldn't have updated post.EditAt when pinning post")

		*post = *saved
	}

	time.Sleep(time.Millisecond * 100)

	post.Message = model.NewId()
	if saved, err := th.App.UpdatePost(post, true); err != nil {
		t.Fatal(err)
	} else if saved.EditAt == post.EditAt {
		t.Fatal("should have updated post.EditAt when updating post message")
	}

	time.Sleep(time.Millisecond * 200)
}

func TestUpdatePostTimeLimit(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	post := &model.Post{}
	*post = *th.BasicPost

	th.App.SetLicense(model.NewTestLicense())

	th.App.UpdateConfig(func(cfg *model.Config) {
		*cfg.ServiceSettings.PostEditTimeLimit = -1
	})
	if _, err := th.App.UpdatePost(post, true); err != nil {
		t.Fatal(err)
	}

	th.App.UpdateConfig(func(cfg *model.Config) {
		*cfg.ServiceSettings.PostEditTimeLimit = 1000000000
	})
	post.Message = model.NewId()
	if _, err := th.App.UpdatePost(post, true); err != nil {
		t.Fatal("should allow you to edit the post")
	}

	th.App.UpdateConfig(func(cfg *model.Config) {
		*cfg.ServiceSettings.PostEditTimeLimit = 1
	})
	post.Message = model.NewId()
	if _, err := th.App.UpdatePost(post, true); err == nil {
		t.Fatal("should fail on update old post")
	}

	th.App.UpdateConfig(func(cfg *model.Config) {
		*cfg.ServiceSettings.PostEditTimeLimit = -1
	})
}

func TestPostReplyToPostWhereRootPosterLeftChannel(t *testing.T) {
	// This test ensures that when replying to a root post made by a user who has since left the channel, the reply
	// post completes successfully. This is a regression test for PLT-6523.
	th := Setup().InitBasic()
	defer th.TearDown()

	channel := th.BasicChannel
	userInChannel := th.BasicUser2
	userNotInChannel := th.BasicUser
	rootPost := th.BasicPost

	if _, err := th.App.AddUserToChannel(userInChannel, channel); err != nil {
		t.Fatal(err)
	}

	if err := th.App.RemoveUserFromChannel(userNotInChannel.Id, "", channel); err != nil {
		t.Fatal(err)
	}

	replyPost := model.Post{
		Message:       "asd",
		ChannelId:     channel.Id,
		RootId:        rootPost.Id,
		ParentId:      rootPost.Id,
		PendingPostId: model.NewId() + ":" + fmt.Sprint(model.GetMillis()),
		UserId:        userInChannel.Id,
		CreateAt:      0,
	}

	if _, err := th.App.CreatePostAsUser(&replyPost, false); err != nil {
		t.Fatal(err)
	}
}

func TestPostChannelMentions(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	channel := th.BasicChannel
	user := th.BasicUser

	channelToMention, err := th.App.CreateChannel(&model.Channel{
		DisplayName: "Mention Test",
		Name:        "mention-test",
		Type:        model.CHANNEL_OPEN,
		TeamId:      th.BasicTeam.Id,
	}, false)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer th.App.PermanentDeleteChannel(channelToMention)

	_, err = th.App.AddUserToChannel(user, channel)
	require.Nil(t, err)

	post := &model.Post{
		Message:       fmt.Sprintf("hello, ~%v!", channelToMention.Name),
		ChannelId:     channel.Id,
		PendingPostId: model.NewId() + ":" + fmt.Sprint(model.GetMillis()),
		UserId:        user.Id,
		CreateAt:      0,
	}

	result, err := th.App.CreatePostAsUser(post, false)
	require.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"mention-test": map[string]interface{}{
			"display_name": "Mention Test",
		},
	}, result.Props["channel_mentions"])

	post.Message = fmt.Sprintf("goodbye, ~%v!", channelToMention.Name)
	result, err = th.App.UpdatePost(post, false)
	require.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"mention-test": map[string]interface{}{
			"display_name": "Mention Test",
		},
	}, result.Props["channel_mentions"])
}

func TestImageProxy(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	th.App.UpdateConfig(func(cfg *model.Config) {
		*cfg.ServiceSettings.SiteURL = "http://mymattermost.com"
	})

	for name, tc := range map[string]struct {
		ProxyType       string
		ProxyURL        string
		ProxyOptions    string
		ImageURL        string
		ProxiedImageURL string
	}{
		"atmos/camo": {
			ProxyType:       model.IMAGE_PROXY_TYPE_ATMOS_CAMO,
			ProxyURL:        "https://127.0.0.1",
			ProxyOptions:    "foo",
			ImageURL:        "http://mydomain.com/myimage",
			ProxiedImageURL: "https://127.0.0.1/f8dace906d23689e8d5b12c3cefbedbf7b9b72f5/687474703a2f2f6d79646f6d61696e2e636f6d2f6d79696d616765",
		},
		"atmos/camo_SameSite": {
			ProxyType:       model.IMAGE_PROXY_TYPE_ATMOS_CAMO,
			ProxyURL:        "https://127.0.0.1",
			ProxyOptions:    "foo",
			ImageURL:        "http://mymattermost.com/myimage",
			ProxiedImageURL: "http://mymattermost.com/myimage",
		},
		"atmos/camo_PathOnly": {
			ProxyType:       model.IMAGE_PROXY_TYPE_ATMOS_CAMO,
			ProxyURL:        "https://127.0.0.1",
			ProxyOptions:    "foo",
			ImageURL:        "/myimage",
			ProxiedImageURL: "/myimage",
		},
		"atmos/camo_EmptyImageURL": {
			ProxyType:       model.IMAGE_PROXY_TYPE_ATMOS_CAMO,
			ProxyURL:        "https://127.0.0.1",
			ProxyOptions:    "foo",
			ImageURL:        "",
			ProxiedImageURL: "",
		},
		"local": {
			ProxyType:       model.IMAGE_PROXY_TYPE_LOCAL,
			ImageURL:        "http://mydomain.com/myimage",
			ProxiedImageURL: "http://mymattermost.com/api/v4/image?url=http%3A%2F%2Fmydomain.com%2Fmyimage",
		},
		"local_SameSite": {
			ProxyType:       model.IMAGE_PROXY_TYPE_LOCAL,
			ImageURL:        "http://mymattermost.com/myimage",
			ProxiedImageURL: "http://mymattermost.com/myimage",
		},
		"local_PathOnly": {
			ProxyType:       model.IMAGE_PROXY_TYPE_LOCAL,
			ImageURL:        "/myimage",
			ProxiedImageURL: "/myimage",
		},
		"local_EmptyImageURL": {
			ProxyType:       model.IMAGE_PROXY_TYPE_LOCAL,
			ImageURL:        "",
			ProxiedImageURL: "",
		},
	} {
		t.Run(name, func(t *testing.T) {
			th.App.UpdateConfig(func(cfg *model.Config) {
				cfg.ImageProxySettings.Enable = model.NewBool(true)
				cfg.ImageProxySettings.ImageProxyType = model.NewString(tc.ProxyType)
				cfg.ImageProxySettings.RemoteImageProxyOptions = model.NewString(tc.ProxyOptions)
				cfg.ImageProxySettings.RemoteImageProxyURL = model.NewString(tc.ProxyURL)
			})

			post := &model.Post{
				Id:      model.NewId(),
				Message: "![foo](" + tc.ImageURL + ")",
			}

			list := model.NewPostList()
			list.Posts[post.Id] = post

			assert.Equal(t, "![foo]("+tc.ProxiedImageURL+")", th.App.PostWithProxyAddedToImageURLs(post).Message)

			assert.Equal(t, "![foo]("+tc.ImageURL+")", th.App.PostWithProxyRemovedFromImageURLs(post).Message)
			post.Message = "![foo](" + tc.ProxiedImageURL + ")"
			assert.Equal(t, "![foo]("+tc.ImageURL+")", th.App.PostWithProxyRemovedFromImageURLs(post).Message)

			if tc.ImageURL != "" {
				post.Message = "![foo](" + tc.ImageURL + " =500x200)"
				assert.Equal(t, "![foo]("+tc.ProxiedImageURL+" =500x200)", th.App.PostWithProxyAddedToImageURLs(post).Message)
				assert.Equal(t, "![foo]("+tc.ImageURL+" =500x200)", th.App.PostWithProxyRemovedFromImageURLs(post).Message)
				post.Message = "![foo](" + tc.ProxiedImageURL + " =500x200)"
				assert.Equal(t, "![foo]("+tc.ImageURL+" =500x200)", th.App.PostWithProxyRemovedFromImageURLs(post).Message)
			}
		})
	}
}

func TestMaxPostSize(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Description         string
		StoreMaxPostSize    int
		ExpectedMaxPostSize int
		ExpectedError       *model.AppError
	}{
		{
			"error fetching max post size",
			0,
			model.POST_MESSAGE_MAX_RUNES_V1,
			model.NewAppError("TestMaxPostSize", "this is an error", nil, "", http.StatusBadRequest),
		},
		{
			"4000 rune limit",
			4000,
			4000,
			nil,
		},
		{
			"16383 rune limit",
			16383,
			16383,
			nil,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.Description, func(t *testing.T) {
			t.Parallel()

			mockStore := &storetest.Store{}
			defer mockStore.AssertExpectations(t)

			mockStore.PostStore.On("GetMaxPostSize").Return(
				storetest.NewStoreChannel(store.StoreResult{
					Data: testCase.StoreMaxPostSize,
					Err:  testCase.ExpectedError,
				}),
			)

			app := App{
				Srv: &Server{
					Store:  mockStore,
					config: atomic.Value{},
				},
			}

			assert.Equal(t, testCase.ExpectedMaxPostSize, app.MaxPostSize())
		})
	}
}

func TestDeletePostWithFileAttachments(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	// Create a post with a file attachment.
	teamId := th.BasicTeam.Id
	channelId := th.BasicChannel.Id
	userId := th.BasicUser.Id
	filename := "test"
	data := []byte("abcd")

	info1, err := th.App.DoUploadFile(time.Date(2007, 2, 4, 1, 2, 3, 4, time.Local), teamId, channelId, userId, filename, data)
	if err != nil {
		t.Fatal(err)
	} else {
		defer func() {
			<-th.App.Srv.Store.FileInfo().PermanentDelete(info1.Id)
			th.App.RemoveFile(info1.Path)
		}()
	}

	post := &model.Post{
		Message:       "asd",
		ChannelId:     channelId,
		PendingPostId: model.NewId() + ":" + fmt.Sprint(model.GetMillis()),
		UserId:        userId,
		CreateAt:      0,
		FileIds:       []string{info1.Id},
	}

	post, err = th.App.CreatePost(post, th.BasicChannel, false)
	assert.Nil(t, err)

	// Delete the post.
	post, err = th.App.DeletePost(post.Id, userId)
	assert.Nil(t, err)

	// Wait for the cleanup routine to finish.
	time.Sleep(time.Millisecond * 100)

	// Check that the file can no longer be reached.
	_, err = th.App.GetFileInfo(info1.Id)
	assert.NotNil(t, err)
}

func TestCreatePost(t *testing.T) {
	t.Run("call PreparePostForClient before returning", func(t *testing.T) {
		th := Setup().InitBasic()
		defer th.TearDown()

		th.App.UpdateConfig(func(cfg *model.Config) {
			*cfg.ExperimentalSettings.EnablePostMetadata = false
			*cfg.ImageProxySettings.Enable = true
			*cfg.ImageProxySettings.ImageProxyType = "atmos/camo"
			*cfg.ImageProxySettings.RemoteImageProxyURL = "https://127.0.0.1"
			*cfg.ImageProxySettings.RemoteImageProxyOptions = "foo"
		})

		imageURL := "http://mydomain.com/myimage"
		proxiedImageURL := "https://127.0.0.1/f8dace906d23689e8d5b12c3cefbedbf7b9b72f5/687474703a2f2f6d79646f6d61696e2e636f6d2f6d79696d616765"

		post := &model.Post{
			ChannelId: th.BasicChannel.Id,
			Message:   "![image](" + imageURL + ")",
			UserId:    th.BasicUser.Id,
		}

		rpost, err := th.App.CreatePost(post, th.BasicChannel, false)
		require.Nil(t, err)
		assert.Equal(t, "![image]("+proxiedImageURL+")", rpost.Message)
	})
}

func TestPatchPost(t *testing.T) {
	t.Run("call PreparePostForClient before returning", func(t *testing.T) {
		th := Setup().InitBasic()
		defer th.TearDown()

		th.App.UpdateConfig(func(cfg *model.Config) {
			*cfg.ExperimentalSettings.EnablePostMetadata = false
			*cfg.ImageProxySettings.Enable = true
			*cfg.ImageProxySettings.ImageProxyType = "atmos/camo"
			*cfg.ImageProxySettings.RemoteImageProxyURL = "https://127.0.0.1"
			*cfg.ImageProxySettings.RemoteImageProxyOptions = "foo"
		})

		imageURL := "http://mydomain.com/myimage"
		proxiedImageURL := "https://127.0.0.1/f8dace906d23689e8d5b12c3cefbedbf7b9b72f5/687474703a2f2f6d79646f6d61696e2e636f6d2f6d79696d616765"

		post := &model.Post{
			ChannelId: th.BasicChannel.Id,
			Message:   "![image](http://mydomain/anotherimage)",
			UserId:    th.BasicUser.Id,
		}

		rpost, err := th.App.CreatePost(post, th.BasicChannel, false)
		require.Nil(t, err)
		assert.NotEqual(t, "![image]("+proxiedImageURL+")", rpost.Message)

		patch := &model.PostPatch{
			Message: model.NewString("![image](" + imageURL + ")"),
		}

		rpost, err = th.App.PatchPost(rpost.Id, patch)
		require.Nil(t, err)
		assert.Equal(t, "![image]("+proxiedImageURL+")", rpost.Message)
	})
}

func TestUpdatePost(t *testing.T) {
	t.Run("call PreparePostForClient before returning", func(t *testing.T) {
		th := Setup().InitBasic()
		defer th.TearDown()

		th.App.UpdateConfig(func(cfg *model.Config) {
			*cfg.ExperimentalSettings.EnablePostMetadata = false
			*cfg.ImageProxySettings.Enable = true
			*cfg.ImageProxySettings.ImageProxyType = "atmos/camo"
			*cfg.ImageProxySettings.RemoteImageProxyURL = "https://127.0.0.1"
			*cfg.ImageProxySettings.RemoteImageProxyOptions = "foo"
		})

		imageURL := "http://mydomain.com/myimage"
		proxiedImageURL := "https://127.0.0.1/f8dace906d23689e8d5b12c3cefbedbf7b9b72f5/687474703a2f2f6d79646f6d61696e2e636f6d2f6d79696d616765"

		post := &model.Post{
			ChannelId: th.BasicChannel.Id,
			Message:   "![image](http://mydomain/anotherimage)",
			UserId:    th.BasicUser.Id,
		}

		rpost, err := th.App.CreatePost(post, th.BasicChannel, false)
		require.Nil(t, err)
		assert.NotEqual(t, "![image]("+proxiedImageURL+")", rpost.Message)

		post.Id = rpost.Id
		post.Message = "![image](" + imageURL + ")"

		rpost, err = th.App.UpdatePost(post, false)
		require.Nil(t, err)
		assert.Equal(t, "![image]("+proxiedImageURL+")", rpost.Message)
	})
}
