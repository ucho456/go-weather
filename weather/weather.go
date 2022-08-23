package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Area     string `json:"targetArea"`
	HeadLine string `json:"headlineText"`
	Body     string `json:"text"`
}

func (w *Weather) ToString() string {
	area := fmt.Sprintf("%sの天気です。\n", w.Area)
	head := fmt.Sprintf("%s\n", w.HeadLine)
	body := fmt.Sprintf("%s\n", w.Body)
	result := area + head + body
	return result
}

func GetWeather(c string) (string, error) {
	url := fmt.Sprintf("https://www.jma.go.jp/bosai/forecast/data/overview_forecast/%s.json", c)
	body, err := reqWeather(url)
	if err != nil {
		return "", err
	}

	weather, err := toStruct(body)
	if err != nil {
		return "", err
	}

	result := weather.ToString()
	return result, nil
}

func reqWeather(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func toStruct(body []byte) (*Weather, error) {
	weather := new(Weather)
	if err := json.Unmarshal(body, weather); err != nil {
		return nil, err
	}
	return weather, nil
}
