package subscriber

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Subscriber interface {
	Subscribe(topic string, qod byte, handler func(client mqtt.Client, message mqtt.Message)) error
}
