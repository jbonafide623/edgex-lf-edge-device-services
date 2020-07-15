package device

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gitlab.com/bonafide-technology/edgex-lf-device-services/internal/mqtt/publisher"
)

type Device interface {
	PublishData(topic string, p publisher.Publisher)

	PublishResponse(uuid string, topic string, p publisher.Publisher)

	Update(message mqtt.Message)
}
