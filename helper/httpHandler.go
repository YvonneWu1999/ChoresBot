package helper

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewHttpHandler(bot *linebot.Client) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			LogError("Bot Parse Error", err)
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		// random reply message to keyword "打掃完畢😎"
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					replyChores := [5]string{"好棒棒 🍭", "讚讚 👍", "good job 🙌", "yes queen 🫶", "nice 🤙"}
					if message.Text == "打掃完畢😎" {
						LogInfo(fmt.Sprintf("groupID:%s", event.Source.GroupID))
						profile, err := bot.GetProfile(event.Source.UserID).Do()
						if err != nil {
							LogError("Get Profile Error", err)
							return
						}
						replyMessage := profile.DisplayName + " " + replyChores[rand.Intn(len(replyChores))]
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
							LogError("Send Reply Done Messege Error", err)
						}
					}
				case *linebot.StickerMessage:
					// replyMessage := fmt.Sprintf(
					// 	"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					// Seed the random number generator with the current time
					rand.Seed(time.Now().UnixNano())

					// Generate a random number between 52114110 and 52114149
					randomNumber := rand.Intn(52114149-52114110+1) + 52114110
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("11539", fmt.Sprintf("%v", randomNumber))).Do(); err != nil {
						LogError("Send Reply Sticker Messege Error", err)
					}
				}
			}
		}
	}
}
