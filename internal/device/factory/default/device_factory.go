package _default

import (
	"gitlab.com/bonafide-technology/edgex-lf-device-services/internal/device"
	"gitlab.com/bonafide-technology/edgex-lf-device-services/internal/device/ac_fan"
	"gitlab.com/bonafide-technology/edgex-lf-device-services/internal/device/temperature_sensor"
)

type factory struct{}

func New() factory {
	return factory{}
}

func (factory) Get(t string) device.Device {
	if t == "temperature" {
		return temperature_sensor.New()
	} else {
		return ac_fan.New()
	}
}
