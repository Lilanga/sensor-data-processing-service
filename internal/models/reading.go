package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"go.chromium.org/luci/common/data/strpair"
)

type Reading struct {
	Time        time.Time
	SensorId    string
	Temperature float64
	Humidity    float64
}

func GetReadingFromMqttPayload(payload string) (*Reading, error) {
	fields := strings.Split(payload, ",")

	if len(fields) != 3 {
		return nil, errors.New("Invalid payload format")
	}

	values := strpair.ParseMap(fields)
	var temp float64 = 0
	var humidity float64 = 0
	temp, err := strconv.ParseFloat(values.Get("t"), 64)
	if err != nil {
		return nil, err
	}

	humidity, err = strconv.ParseFloat(values.Get("h"), 64)
	if err != nil {
		return nil, err
	}

	reading := &Reading{
		Time:        time.Now(),
		SensorId:    values.Get("n"),
		Temperature: temp,
		Humidity:    humidity,
	}
	return reading, nil
}
