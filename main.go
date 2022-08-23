package GoWeather

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ucho456/go_weather/weather"
)

func GoWeather(w http.ResponseWriter, r *http.Request) {
	client, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	result, err := weather.GetWeather()
	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(result)
	if _, err := client.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
