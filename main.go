package main

import (
	"context"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ucho456/go_weather/weather"
)

func main() {
	GoWeather(context.TODO())
}

func GoWeather(ctx context.Context) error {
	bot, err := linebot.New(
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
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
	return nil
}
