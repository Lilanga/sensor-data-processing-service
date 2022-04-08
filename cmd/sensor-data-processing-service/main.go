package main

import (
	"fmt"
	"sync"

	mqttLib "github.com/Lilanga/sensor-data-processing-service/pkg/mqtt"
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

//todo: still need mqtt library structs here, need to create higher order function or message topic approach
func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Message handler message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
