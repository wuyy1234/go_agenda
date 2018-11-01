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
	"agenda/entity"

	"github.com/spf13/cobra"
)

// searchMeetingCmd represents the searchMeeting command
var searchMeetingCmd = &cobra.Command{
	Use:   "searchMeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sm, _ := cmd.Flags().GetString("startMonth")
		sd, _ := cmd.Flags().GetString("startDay")
		em, _ := cmd.Flags().GetString("endMonth")
		ed, _ := cmd.Flags().GetString("endDay")
		users = entity.READUSERS()
		meetings = entity.READMEETINGS()
		current = entity.GetCurrentUserName()
		currentIndex := -1
		//遍历找到当前用户所在位置
		for i, user := range users {
			if user.Username == current {
				currentIndex = i
				break
			}
		}
		//遍历所有会议
		for i, meeting := range meetings {
			for j, time := range meeting.MeetingTime {
				//在查询期间
				if (time.month > sm || (time.month == sm && time.day >= sd)) && (time.month < em || (time.month == em && time.day <= sd)) {
					var flag = false //判断是否参与会议或发起会议
					//判断是否为发起人
					if meeting.Sponsor == current {
						log.println("You Sponsor: ")
						flag = true
					} else { //判断是否为参与者
						for k, par := range meeting.Participators {
							if par == current {
								log.println("You Participate: ")
								flag = true
							}
						}
					}
					//如果参与了会议，则打印相关信息
					if flag {
						log.println("Title: " + meeting.Title + "Sponsor: " + meeting.Sponsor)
						log.println("Date: " + time.month + "." + time.day)
						for l, tid := range time.timeID {
							switch tid {
							case 1:
								log.println("10:00~11:00")
							case 2:
								log.println("11:00~12:00")
							case 3:
								log.println("15:00~16:00")
							case 4:
								log.println("16:00~17:00")
							}
						}
						log.println("Participate: " + meeting.Participators)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//得到会议名称[-meeting meeting]
	searchMeetingParCmd.Flags().StringP("meeting", "m", "default meeting", "search meeting participants")
}
