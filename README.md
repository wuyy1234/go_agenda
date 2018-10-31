# go_agenda
## 伍宇阳 负责初始化，构建框架以及写register.go login.go logout.go searchUser.go
## 吴啸林 负责user部分函数
## 王永杰 负责meeting部分函数

## Agenda 业务需求

### 用户注册

* 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
* 如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。
* 使用说明：register -u username -P password -e email -p phone
* 运行结果：
```
//成功
[wyy@centos7 agenda]$ go run main.go register -u wuyy -P 2333 -e 121212@163.com -p 17766727
2018/10/31 19:54:39 register called by wuyy 2333 121212@163.com 17766727
2018/10/31 19:54:39 READUSER success
2018/10/31 19:54:39 WRITEUSER success

//已注册，失败
[wyy@centos7 agenda]$ go run main.go register -u wuyy -P 2333 -e 121212@163.com -p 17766727
2018/10/31 20:04:25 register called by wuyy 2333 121212@163.com 17766727
2018/10/31 20:04:25 READUSER success
2018/10/31 20:04:25 username already registered

```

### 用户登录

* 用户使用用户名和密码登录 Agenda 系统。
* 用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。
* 使用说明：login -u username -p password
* 运行结果：
```
//成功
[wyy@centos7 agenda]$ go run main.go login -u wuyy -p 2333
2018/10/31 19:55:49 READUSER success
2018/10/31 19:55:49 user login success

//失败
[wyy@centos7 agenda]$ go run main.go login -u wwyy -p 121212
2018/10/31 19:55:36 READUSER success
2018/10/31 19:55:36 user login failed

```

### 用户登出

* 已登录的用户登出系统后，只能使用用户注册和用户登录功能。
* 使用说明：logout
* 运行结果：
```
[wyy@centos7 agenda]$ go run main.go logout
2018/10/31 19:56:22 logout success

```

### 用户查询

* 已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。
* 使用说明：userSearch
* 运行结果：
```
[wyy@centos7 agenda]$ go run main.go searchUser2018/10/31 19:56:13 READUSER success
2018/10/31 19:56:13 list all the login users' username&email&phone
2018/10/31 19:56:13  username:TestUser email:123@qq.com phone:123
2018/10/31 19:56:13  username:wuyy email:121212@163.com phone:17766727

```

### 用户删除

* 已登录的用户可以删除本用户账户（即销号）。
* 操作成功，需反馈一个成功注销的信息；否则，反馈一个失败注销的信息。
* 删除成功则退出系统登录状态。删除后，该用户账户不再存在。
* 用户账户删除以后：
   * 以该用户为 发起者 的会议将被删除
   * 以该用户为 参与者 的会议将从 参与者 列表中移除该用户。若因此造成会议 参与者 人数为0，则会议也将被删除。
* 使用说明：
* 运行结果：
```

```

### 创建会议

* 已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册 用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：
   * 会议主题(title)（在会议列表中具有唯一性）
   * 会议参与者(participator)
   * 会议起始时间(start time)
   * 会议结束时间(end time)
   * 注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。
   * 用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。
* 使用说明：createMeeting [-meeting meeting] [-Month month] [-day day] [-time time]1/2/3/4 [-command command]update/new 增加新的时间/创建新的会议
* 运行结果：
//成功
```
go run main.go createMeeting  -m funk_xiao_li -M 1 -d 1 -t 1 -c new
```
//更新时间冲突
```
go run main.go createMeeting  -m funk_xiao_li -M 1 -d 1 -t 1 -c update
```
//创建时间冲突
```
go run main.go createMeeting  -m not_funk_xiao_li -M 1 -d 1 -t 1 -c update
```
//创建题目冲突
```
go run main.go createMeeting  -m funk_xiao_li -M 2 -d 1 -t 1 -c update
```


### 增删会议参与者

* 已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。
* 增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。
* 删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。
* 使用说明：changeMeetingPar [-meeting meeting] 指令名称[-command a/d] 用户名称[-par name]
* 运行结果：
//成功增加
```
go run main.go changeMeetingPar -m funk_xiao_li -c a -p TestUser
```
//增加失败，重复增加
```
go run main.go changeMeetingPar -m funk_xiao_li -c a -p TestUser
```

### 查询会议

* 已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。
* 用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。
* 在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。
* 注意，查询会议的结果应包括用户作为 发起者或参与者 的会议。
* 使用说明：searchMeeting [-meeting meeting]
* 运行结果：
```
go run main.go searchMeeting -m funk_xiao_li
```

### 取消会议

* 已登录的用户可以取消 自己发起 的某一会议安排。
* 取消会议时，需提供唯一标识：会议主题（title）。
* 使用说明：deleteMeeting [-meeting meeting]
* 运行结果：
//成功取消
```
go run main.go deleteMeeting -m funk_xiao_li
```
//取消失败
```
go run main.go deleteMeeting -m funk_xiao_li
```

### 退出会议

* 已登录的用户可以退出 自己参与 的某一会议安排。
* 退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议 参与者 人数为0，则会议也将被删除。
* 使用说明：exitMeeting [-meeting meeting]
* 运行结果：
```
go run main.go exitMeeting -m funk_xiao_li
```

### 清空会议

* 已登录的用户可以清空 自己发起 的所有会议安排。
* 使用说明：clearMeeting [-pass password]
* 运行结果：
```
go run main.go clearMeeting -m funk_xiao_li
```


