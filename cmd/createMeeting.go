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
	"fmt"

	"github.com/spf13/cobra"
)

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("meeting")
		month, _ := cmd.Flags().GetString("month")
		day, _ := cmd.Flags().GetString("day")
		time, _ := cmd.Flags().GetString("time")
		com, _ := cmd.Flags().GetString("command")
		users = entity.READUSERS()
		meetings = entity.READMEETINGS()
		current = entity.GetCurrentUserName()
		//添加新的时间
		if com == "update" {
			for i, meeting := range meetings {
				if meeting.Title != title {
					continue
				}
				if meeting.Sponsor != current {
					log.println("Wrong! You aren't the Sponsor!")
					return
				}
				for j, readyTime := range meeting.MeetingTime {
					//增加当天不同的时间段
					if readyTime.day == day && readyTime.month == month {
						for k, tid := range readyTime.timeID {
							//时间重复
							if tid == time {
								log.println("Wrong! You had apply this time")
								return
							}
						}
						readyTime.timeID = append(readyTime, time)
						log.println("Apply Success!")
						return 
					}
					//增加不同的时间
					else {
						newTime := entity.Time{timeID: [...]var{time}, day: day, month: month}
						log.println("Apply Success!")
						return 
					}
				}
			}
			log.println("Wrong! You should use -command new")
		}
		//新建会议事项
		else {
			//名称、时间查重
			for i, meeting := range meetings {
				//名称查重
				if meeting.Title == title {
					log.println("Wrong! You should use -command update")
					return
				}
				//时间查重
				for j, readyTime := range meeting.MeetingTime {
					//同一天不同时间段查重
					if readyTime.day == day && readyTime.month == month {
						for k, tid := range readyTime.timeID {
							//时间重复
							if tid == time {
								log.println("Wrong! Time had applied!")
								return
							}
						}
					}
				}
				//创建新的会议事件并加入会议列表
				newTime := entity.Time{timeID: [...]var{time}, day: day, month: month}
				newMeeting := entity.Meeting{Title: title, Sponsor: current, Paticipators: [...]var{}, MeetingTime: [...]var{newTime} }
				meetings = append(meetings, newMeeting)
				//为操作者添加会议事件
				for j, user := range users {
					if user.Username == current {
						user.SponsorMeeting = append(user.SponsorMeeting, title)
						break
					}
				}
				log.println("Apply Success! Please add Paticipators!")
				return 
			}
		}
		//记录写回
		entity.WRITEUSER(users)
		entity.WRITEMEETINGS(meetings)
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//得到会议名称[-meeting meeting] [-Month month] [-Day day] [-time time]1/2/3/4 [-command command]update/new 增加新的时间/创建新的会议
	createMeetingParCmd.Flags().StringP("meeting", "m", "default meeting", "create meeting participants")
	createMeetingParCmd.Flags().IntP("month", "M", 1, "create month participants")
	createMeetingParCmd.Flags().IntP("day", "d", 1, "create day participants")
	createMeetingParCmd.Flags().IntP("time", "t", 1, "create time participants")
	createMeetingParCmd.Flags().StringP("command", "c", "a", "create command participants")
}
