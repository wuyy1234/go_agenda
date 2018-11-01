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
	"log"
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
		sm, _ := cmd.Flags().GetInt("startMonth")
		sd, _ := cmd.Flags().GetInt("startDay")
		em, _ := cmd.Flags().GetInt("endMonth")
		ed, _ := cmd.Flags().GetInt("endDay")
		meetings := entity.READMEETINGS()
		current := entity.GetCurrentUserName()
		//遍历所有会议
		for _, meeting := range meetings {
			for _, time := range meeting.MeetingTime {
				//在查询期间
				if (time.Month > sm || (time.Month == sm && time.Day >= sd)) && (time.Month < em || (time.Month == em && time.Day <= ed)) {
					var flag = false //判断是否参与会议或发起会议
					//判断是否为发起人
					if meeting.Sponsor == current {
						log.Println("You Sponsor: ")
						flag = true
					} else { //判断是否为参与者
						for _, par := range meeting.Participators {
							if par == current {
								log.Println("You Participate: ")
								flag = true
							}
						}
					}
					//如果参与了会议，则打印相关信息
					if flag {
						log.Println("Title: " + meeting.Title + "Sponsor: " + meeting.Sponsor)
						log.Print("Date: ")
						log.Print(time.Month)
						log.Print(".")
						log.Println(time.Day)
						for _, tid := range time.TimeID {
							switch {
							case tid == 1:
								log.Println("10:00~11:00")
							case tid == 2:
								log.Println("11:00~12:00")
							case tid == 3:
								log.Println("16:00~17:00")
							default:
								log.Println("17:00~18:00")
							}
						}
						log.Print("Participate: ")
						log.Println(meeting.Participators)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchMeetingCmd)
	//得到会议名称[-meeting meeting]
	searchMeetingCmd.Flags().IntP("startMonth", "sm", 1, "Start Month")
	searchMeetingCmd.Flags().IntP("startDay", "sd", 1, "Start Day")
	searchMeetingCmd.Flags().IntP("endMonth", "em", 1, "End Month")
	searchMeetingCmd.Flags().IntP("endDay", "ed", 1, "End Day")
}
