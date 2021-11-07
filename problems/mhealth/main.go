package main

import (
	"fmt"
	"strings"
	"time"
)

type notification struct {
	//	name        string
	// occurence int
	counter  int
	duration time.Time
	//waitTime    time.Time
	lastupdated time.Time
}

type user string

var users []user

// type notifications_type string

// const (
// 	info     notifications_type = "Info"
// 	blocker  notifications_type = "Blocker"
// 	critical notifications_type = "Critical"
// 	warning  notifications_type = "Warninig"
// )

func ProcessNotification() {

}

func checknontificationcondition(noti string, paramaeters notification) (notification, bool) {
	time_diff := time.Since(paramaeters.duration)
	if time_diff.Seconds() >= float64(100*time.Second) {
		paramaeters.lastupdated = time.Now()
		if paramaeters.counter > 10 {
			paramaeters.counter = 0
			return paramaeters, true
		}
	} else {
		paramaeters.counter = +1
	}
	return paramaeters, false
}

func checkType(notification_type string, subs_map map[string]notification) {
	notify_users := subs_map[notification_type]
	updated_parameter, notification_Status := checknontificationcondition(notification_type, notify_users)
	if notification_Status {
		fmt.Println(notification_Status)
		for _, i := range users {
			fmt.Println(i)
		}
	}
	subs_map[notification_type] = updated_parameter
}

func main() {
	input := []string{"2019-01-07 14:52:33 Warning data",
		"2019-01-07 14:52:34 Critical data",
		"2019-01-07 14:52:35 Info data",
		"2019-01-07 14:52:36 Critical data",
		"2019-01-07 14:52:37 Critical data",
		"2019-01-07 14:52:38 Critical data",
		"2019-01-07 14:52:39 Critical data",
		"2019-01-07 14:52:40 Critical data",
		"2019-01-07 14:52:41 Warning data",
		"2019-01-07 14:52:42 Critical data",
		"2019-01-07 14:52:43 Warning data",
		"2019-01-07 14:52:44 Critical data"}
	for i := 0; i < 5; i++ {
		temp := user(fmt.Sprintf(`user%d`, i))
		users = append(users, temp)
	}
	fmt.Println(users)
	subscription := make(map[string]notification)

	for _, i := range input {
		temp := strings.Split(i, " ")
		checkType(temp[2], subscription)
	}

}
