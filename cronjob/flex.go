package main

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func NewChoreMessage(weekNum int, people []string, assignments map[string]string) *linebot.BubbleContainer {
	choreTaskBox := NewChoreTaskBox(people, assignments)
	contents := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   fmt.Sprintf("Week%v打掃提醒🔔", weekNum),
					Weight: "bold",
					Size:   "xxl",
					Color:  "#0000ff",
				},
				choreTaskBox,
			},
		},
	}
	return contents
}

func NewChoreTaskBox(people []string, assignments map[string]string) *linebot.BoxComponent {
	taskContent := []linebot.FlexComponent{}
	for name, assignment := range assignments {
		if assignment == "" {
			assignment = "不用做家事"
		}

		choreRightComponentContent := []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:      linebot.FlexComponentTypeText,
				Text:      fmt.Sprintf("%s", name),
				OffsetTop: "8px",
				Weight:    "bold",
			},
			&linebot.TextComponent{
				Type:      linebot.FlexComponentTypeText,
				Text:      fmt.Sprintf("%s", assignment),
				OffsetTop: "8px",
				Weight:    "bold",
			},
		}

		if assignment != "不用做家事" {
			choreRightComponentContent = append(choreRightComponentContent, &linebot.ButtonComponent{
				Type:      linebot.FlexComponentTypeButton,
				OffsetTop: "10px",
				Style:     "primary",
				Color:     "#00b900",
				Action: &linebot.MessageAction{
					Label: "Done",
					Text:  "打掃完畢😎",
				},
			})
		}

		taskContent = append(taskContent,
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.ImageComponent{
						Type:        linebot.FlexComponentTypeImage,
						URL:         viper.GetString(fmt.Sprintf("PICURL%v", indexOf(name, people)+1)),
						Align:       "start",
						OffsetTop:   "8px",
						AspectRatio: "1:1",
					}, &linebot.BoxComponent{
						Type:     linebot.FlexComponentTypeBox,
						Layout:   linebot.FlexBoxLayoutTypeVertical,
						Contents: choreRightComponentContent,
					},
				},
			},
		)
	}

	return &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: taskContent,
	}
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
