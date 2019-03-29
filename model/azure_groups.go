// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

var ValidAzureGroups = map[string]bool{
	"ALL":        true,
	"STUDENTS":   true,
	"PARENTS":    true,
	"TEACHERS":   true,
	"MODERATORS": true,
}

var GroupsMapping = map[string]string{
	"All_Staff":      "TEACHERS",
	"SA_ADMINSTAFF":  "MODERATORS",
	"Parent_CC":      "PARENTS",
	"Parent_CC_MC":   "PARENTS",
	"Parent_MC":      "PARENTS",
	"Parent_WHJ":     "PARENTS",
	"Parent_WHS":     "PARENTS",
	"Parent_WHS_WHJ": "PARENTS",
	"All_Parents":    "PARENTS",
	"All_Students":   "STUDENTS",
}
