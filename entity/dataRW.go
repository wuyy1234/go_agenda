package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*User 用户*/
type User struct {
	Username           string
	Password           string
	Email              string
	Phone              string
	SponsorMeeting     []string //发起的会议title
	ParticipateMeeting []string //参与的会议title
}

/*Time 时间*/
type Time struct {
	Hour   int
	Minute int
}

/*Meeting 会议*/
type Meeting struct {
	Title        string
	Sponsor      string   //发起者
	Paticipators []string //参与者用户名
	StartTime    Time
	EndTime      Time
}

/*READUSERS  读取文件*/
func READUSERS() (user []User) {
	dir, err := os.Getwd()
	checkerr(err)
	b, err := ioutil.ReadFile(dir + "/entity/Users.txt")
	checkerr(err)

	//json转变为对象
	var users []User
	json.Unmarshal(b, &users)

	log.Println("READUSER success")
	return users
}

/*WRITEUSER 写入文件*/
func WRITEUSER(users []User) {
	dir, err := os.Getwd()
	checkerr(err)
	data, err := json.Marshal(users)
	checkerr(err)
	b := []byte(data)
	err = ioutil.WriteFile(dir+"/entity/Users.txt", b, 0777)
	checkerr(err)
	log.Println("WRITEUSER success")
}

/*READMEETINGS 读取会议信息*/
func READMEETINGS() (meetings []Meeting) {
	dir, err := os.Getwd()
	checkerr(err)
	b, err := ioutil.ReadFile(dir + "/entity/Meetings.txt")
	checkerr(err)
	//json转变为对象
	json.Unmarshal(b, &meetings)
	log.Println("READMEETINGS success")
	return meetings
}

/*WRITEMEETINGS 写入会议信息*/
func WRITEMEETINGS(meetings []Meeting) {
	dir, err := os.Getwd()
	checkerr(err)
	data, err := json.Marshal(meetings)
	checkerr(err)
	b := []byte(data)
	err = ioutil.WriteFile(dir+"/entity/Meetings.txt", b, 0777)
	checkerr(err)
	log.Println("WRITEMEETINGS success")
}

/*READCURRENTUSER 读取当前登陆的用户信息*/
func READCURRENTUSER() (user User) {
	dir, err := os.Getwd()
	checkerr(err)
	b, err := ioutil.ReadFile(dir + "/entity/CurrentUser.txt")
	checkerr(err)
	//json转变为对象
	var users User
	json.Unmarshal(b, &users)
	log.Println("READCURRENTUSER success")
	return users
}

/*WRITECURRENTUSER 编辑当前登陆的用户信息*/
func WRITECURRENTUSER(users User) {
	dir, err := os.Getwd()
	checkerr(err)
	data, err := json.Marshal(users)
	checkerr(err)
	b := []byte(data)
	err = ioutil.WriteFile(dir+"/entity/CurrentUser.txt", b, 0777)
	checkerr(err)
	log.Println("WRITECURRENTUSER success")
}

func checkerr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}
