package factory

import "gitlab.com/bonafide-technology/edgex-lf-device-services/internal/device"

type DeviceFactory interface {
	Get(t string) device.Device
}
