// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

// clearMeetingCmd represents the clearMeeting command
var clearMeetingCmd = &cobra.Command{
	Use:   "clearMeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_password_, _ := cmd.Flags().GetString("pass")
		users = entity.READUSERS()
		meetings = entity.READMEETINGS()
		current = entity.GetCurrentUserName()
		for i, user := range users {
			if user.Username != current {
				continue
			}
			//密码不正确
			if user.Password != _password_ {
				log.println("Wrong password!")
				return
			}
			myClearMeeting()
		}
		//记录写回
		entity.WRITEUSER(users)
		entity.WRITEMEETINGS(meetings)
		return
	},
}

func myClearMeeting () {
	for i, user := range users {
		//删除主持的所有会议
		for j, title := range user.SponsorMeeting {
			entity.myDeleteMeeting(title)
		}
		//退出参加的所有会议
		for j, title := range user.ParticipateMeeting {
			entity.myExitMeeting(title)
		}
	}
}

func init() {
	rootCmd.AddCommand(clearMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//得到用户密码[-pass password]
	clearMeetingCmd.Flags().StringP("pass", "p", "", "clear meeting")
}
