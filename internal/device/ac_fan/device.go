package ac_fan

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gitlab.com/bonafide-technology/edgex-lf-device-services/internal/mqtt/publisher"
)

type device struct {
	name  string
	state string
}

func New() *device {
	return &device{
		name:  "ACFan",
		state: "off",
	}
}

func (d *device) PublishData(topic string, p publisher.Publisher) {
	type data struct {
		Name  string `json:"name"`
		State string `json:"state"`
		Cmd   string `json:"cmd"`
	}
	b, _ := json.Marshal(&data{
		Name:  d.name,
		State: d.state,
		Cmd:   "state",
	})

	if err := p.Publish(topic, 0, false, b); err != nil {
		fmt.Println("Publish ", topic, "error: ", err.Error())
	}
}

func (d *device) PublishResponse(uuid string, topic string, p publisher.Publisher) {
	type response struct {
		Uuid  string `json:"uuid"`
		State string `json:"state"`
	}

	b, _ := json.Marshal(&response{
		Uuid:  uuid,
		State: d.state,
	})

	if err := p.Publish(topic, 0, false, b); err != nil {
		fmt.Println("Publish ", topic, "error: ", err.Error())
	}
}

func (d *device) Update(m mqtt.Message) {
	type message struct {
		Cmd    string `json:"cmd"`
		Method string `json:"method"`
		Uuid   string `json:"uuid"`
		State  string `json:"state"`
	}

	var inbound *message
	if err := json.Unmarshal(m.Payload(), &inbound); err != nil {
		fmt.Println(err.Error())
	}

	d.state = inbound.State
}
