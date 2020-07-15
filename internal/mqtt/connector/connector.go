package connector

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Connector interface {
	Connect(client mqtt.Client) error
}
