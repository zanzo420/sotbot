package handler

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vascocosta/owm"
)

func GetCurrentWeather(o *owm.Client, loc string) (string, error) {
	l := log.WithFields(log.Fields{
		"action": "handler.GetCurrentWeather",
	})

	l.Debugf("Trying to fetch weather conditions for: %v", loc)
	curWeather, err := o.WeatherByName(loc, "metric")
	if err != nil {
		l.Errorf("Failed to look up weather data: %v", err)
		return "", err
	}

	responseMsg := fmt.Sprintf("The current weather condition in %v, %v is: %v at %.1f°C "+
		"(Min: %.0f°C, Max: %.0f°C), Humidity: %d%%, Air pressure: %.0fhPa.",
		curWeather.Name, curWeather.Sys.Country, curWeather.Weather[0].Description, curWeather.Main.Temp,
		curWeather.Main.TempMin, curWeather.Main.TempMax, curWeather.Main.Humidity, curWeather.Main.Pressure)
	return responseMsg, nil
}
