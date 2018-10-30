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

// deleteMeetingCmd represents the deleteMeeting command
var deleteMeetingCmd = &cobra.Command{
	Use:   "deleteMeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_meeting, _ := cmd.Flags().GetString("meeting")
		users = entity.READUSERS()
		meetings = entity.READMEETINGS()
		current = entity.CurrentUserName
		meetingSize = len(meetings)
		for i := 0; i < meetingSize; i++ {
			if (meetings[i].Title == _meeting_) {
				//判断是否是会议发起人
				if meetings[i].Sponsor == current {
					//删除所有与会人的会议记录
					parSize = len(meetings[i].Paticipators)
					for j := 0; j < parSize; j++ {
						par = meetings[i].Paticipators[j]
						userSize = len(users)
						for k := 0; k < userSize; k++ {
							if users[k].Username == par {
								//删除该与会人的会议记录
								parMeetingSize = len(users[k].ParticipateMeeting)
								for f := 0; f < parMeetingSize; f++ {
									if users[k].ParticipateMeeting[f] == _meeting_ {
										users[k].ParticipateMeeting = append(users[k].ParticipateMeeting[:f], users[k].ParticipateMeeting[f+1:]...)
									}
								}
							}
						}
					}
					//删除发起人的列表记录
					userSize = len(users)
					currentIndex := -1
					for j := 0; j < userSize; j++ {
						if users[j].Username == current {
							currentIndex = j
						}
					}
					sponsorMeetingSize = len(users[currentIndex].SponsorMeeting)
					for j := 0; j < sponsorMeetingSize; j++ {
						if users[currentIndex].SponsorMeeting[j] == _meeting_ {
							users[currentIndex].SponsorMeeting = append(users[currentIndex].SponsorMeeting[:j], users[currentIndex].SponsorMeeting[j+1:]...)
						}
					}
					//删除会议
					meetings = append(meetings[:i], meetings[i+1:]...)
					log.println("Delete Meeting Success!")
					entity.WRITEUSER(users)
					entity.WRITEMEETINGS(meetings)
					return 
				}
			}
		}
		log.println("Dont have this Meeting")
		return 
	},
}

func init() {
	rootCmd.AddCommand(deleteMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//得到会议名称[-meeting meeting]
	// deleteMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
