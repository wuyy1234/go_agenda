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

// changeMeetingParCmd represents the changeMeetingPar command
var changeMeetingParCmd = &cobra.Command{
	Use:   "changeMeetingPar",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_meeting, _ := cmd.Flags().GetString("meeting")
		_par_, _ := cmd.Flags().GetString("par")
		_command_, _ := cmd.Flags().GetString("command")
		users = WRITEUSER()
		meetings = READMEETINGS()
		current = entity.CurrentUserName
		meetingSize = len(meetings)
		userSize = len(users)
		parIndex := -1
		for i := 0; i < userSize; i++ {
			if users[i].Username == _par_ {
				parIndex = i
			}
		}
		if parIndex == -1 {
			log.println("Dont have user named " + _par_)
			return 
		}
		for i := 0; i < meetingSize; i++ {
			if (meetings[i].Title == _meeting_) {
				if meetings[i].Sponsor == _users_ {
					if _command_ == "a" {
						parSize = len(meetings[i].Paticipators)
						for j := 0; j < parSize; j++ {
							if meetings[i].Paticipators[j] == _par_ {
								meetings[i].Paticipators = append(meetings[i].Paticipators[:j], meetings[i].Paticipators[j+1:]...)
								parMeetingSize = len(users[parIndex].ParticipateMeeting)
								for k := 0; k < parMeetingSize; k++ {
									if users[parIndex].ParticipateMeeting[k] == _meeting_ {
										users[parIndex].ParticipateMeeting = append(users[parIndex].ParticipateMeeting[:k], users[parIndex].ParticipateMeeting[k+1:]...)
									}
								}
								log.println("Delete success!")
								return 
							}
						}
						log.println("Dont have particapator name " + _par_);
					} else {
						parSize = len(meetings[i].Paticipators)
						for j := 0; j < parSize; j++ {
							if meetings[i].Paticipators[j] == _par_ {
								return 
							}
						}
						meetings[i].Paticipators = append(meetings[i].Paticipators, _par_)
						log.println("Add success!")
					}
				} else {
					log.println("Dont have privilege!")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(changeMeetingParCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changeMeetingParCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changeMeetingParCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//得到会议名称[-meeting meeting] 指令名称[-command a/d] 用户名称[-par name]
	changeMeetingParCmd.Flags().StringP("meeting", "m", "default meeting", "change meeting participants")
	
}
