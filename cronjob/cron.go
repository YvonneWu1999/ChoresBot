package main

import (
	"ChoresBot/helper"
	"fmt"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func main() {
	// Get env
	helper.NewViper()

	// Init line bot
	bot, err := linebot.New(viper.GetString("CHANNELSECRET"), viper.GetString("CHANNELTOKEN"))
	if err != nil {
		helper.LogError("Init Bot Error", err)
	}
	// Assign chores based on current week
	var currentTime time.Time = time.Now()
	var year, weekNum = currentTime.ISOWeek()
	assignments := AssignChores(weekNum, year, []string{"Aria", "Sarah", "Yvonne"}, []string{"倒垃圾♻️", "吸地🧹", "澆水🪴"})

	// Send chores notification to group
	var messages []linebot.SendingMessage
	contents := NewChoreMessage(weekNum, []string{"Aria", "Sarah", "Yvonne"}, assignments)
	m := linebot.NewFlexMessage(fmt.Sprintf("Week%v打掃提醒🔔", weekNum), contents)
	messages = append(messages, m)
	_, err = bot.PushMessage(viper.GetString("GROUPID"), messages...).Do()
	if err != nil {
		helper.LogError("Send Chores Notify Error", err)
		return
	}
}
