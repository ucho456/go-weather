package GoWeather

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ucho456/go_weather/prefecture"
	"github.com/ucho456/go_weather/weather"
)

func GoWeather(w http.ResponseWriter, r *http.Request) {
	client, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		http.Error(w, "Error init client", http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	events, err := client.ParseRequest(r)
	if err != nil {
		http.Error(w, "Error parse request", http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyMessage := message.Text
				code, err := prefecture.GetPrefectureCode(replyMessage)
				if err != nil {
					resMessage := linebot.NewTextMessage("「東京」または「大阪」のどちらかを入力して下さい。")
					if _, err := client.ReplyMessage(event.ReplyToken, resMessage).Do(); err != nil {
						log.Print(err)
					}
					return
				}

				weatherInfo, err := weather.GetWeather(code)
				if err != nil {
					http.Error(w, "Error request weather API", http.StatusBadRequest)
					log.Fatal(err)
					return
				}

				resMessage := linebot.NewTextMessage(weatherInfo)
				if _, err := client.ReplyMessage(event.ReplyToken, resMessage).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
