package main

import (
	"github.com/shawntoffel/darksky"
	"log"
)

func main(key string, lat float, long float) {
	client := darksky.New(key)
	request := darksky.ForecastRequest{}

	request.Latitude = lat
	request.Longitude = long

	request.Options = darksky.ForecastRequestOptions{Exclude: "hourly,minutely"}

	forecast, err := client.Forecast(request)
	if err != nil {
		log.Fatal("Forecast request error", err)
	}

	Emoji := ""

	switch forecast.Currently.Icon {
	case "clear-day":
		Emoji = "☀️"
	case "partly-cloudy-day":
	case "partly-cloudy-night":
		Emoji = "🌤️"
	case "cloudy":
		Emoji = "☁️"
	case "rain":
		Emoji = "🌧️"
	case "snow":
	case "sleet":
		Emoji = "❄️"
	case "wind":
		Emoji = "🌬️"
	case "fog":
		Emoji = "🌫️"
	default:
		Emoji = "☀️"
	}

	return Emoji
}
