package mqtt

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	client mqtt.Client
	config mqtt.ClientOptions
}

func GetMqttClient(clientID string) *MqttClient {
	var client = &MqttClient{}
	client.init(clientID)
	return client
}

func initClientOptions(clientID string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	host := os.Getenv("MQTT_HOST")
	port := os.Getenv("MQTT_PORT")
	broker := fmt.Sprintf("%s:%s", host, port)
	mqttUser := os.Getenv("MQTT_USER")
	mqttPass := os.Getenv("MQTT_PASS")
	mqttClientPrefix := fmt.Sprintf("%s%s", os.Getenv("MQTT_CLIENT_ID"), clientID)

	opts.AddBroker(broker)
	opts.SetClientID(mqttClientPrefix)
	opts.SetUsername(mqttUser)
	opts.SetPassword(mqttPass)
	opts.SetDefaultPublishHandler(defaultPublishedMessageHandler)
	opts.OnConnect = onConnectHandler
	opts.OnConnectionLost = onConnectionLostHandler

	return opts
}
func (m *MqttClient) init(clientID string) (bool, error) {
	client := mqtt.NewClient(initClientOptions(clientID))

	var err error = nil
	status := true

	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		status = false
		err = token.Error()
	}

	m.client = client
	return status, err
}

func (m *MqttClient) Publish(topic string, qos byte, payload interface{}) mqtt.Token {

	return m.client.Publish(topic, qos, false, payload)
}

func (m *MqttClient) Connect() (bool, error) {

	var err error = nil
	status := true

	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		status = false
		err = token.Error()
	}

	return status, err
}

func (m *MqttClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	return m.client.Subscribe(topic, qos, callback)
}

func defaultPublishedMessageHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Default handler received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func onConnectHandler(client mqtt.Client) {
	log.Println("Client connected successfully")
}

func onConnectionLostHandler(client mqtt.Client, err error) {
	log.Printf("Client connection lost: %v", err)
}
