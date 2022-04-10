package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/Lilanga/sensor-data-processing-service/internal/models"
	mqttLib "github.com/Lilanga/sensor-data-processing-service/pkg/mqtt"
	TimescaleDBClient "github.com/Lilanga/sensor-data-processing-service/pkg/timescaledb"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.env")

	var wg sync.WaitGroup
	wg.Add(1)

	subscribeForTopic("testClient", "esp32/message", wg)

	wg.Wait()
}

func subscribeForTopic(clientID string, topic string, wait sync.WaitGroup) {
	client := mqttLib.GetMqttClient(clientID)
	client.Subscribe(topic, 1, messagePubHandler)

	if token := client.Subscribe("esp32/message", 1, messagePubHandler); token.Wait() && token.Error() != nil {
		wait.Done()
	}
}

// todo: still need mqtt library structs here, need to create higher order function or message topic approach
func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	reading, err := models.GetReadingFromMqttPayload(string(msg.Payload()))
	if err != nil {
		log.Printf("Unable to parse mqtt payload %v\n", err)
		return
	}

	dbClient := TimescaleDBClient.GetTimescaleDBClient()
	err = dbClient.InsertReading(*reading)
	if err != nil {
		log.Printf("Unable to insert mqtt payload %v\n", err)
		return
	}

	fmt.Printf("Message handler inserted message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
