package provider

import (
	"strings"

	"github.com/denverdino/aliyungo/ecs"
)

const (
	DevicePrefixLocal = "/dev/vd"
	DevicePrefixAPI   = "/dev/xvd"
)

// Status of disks
type DiskStatus string

const (
	DiskStatusInUse     = DiskStatus("In_use")
	DiskStatusAvailable = DiskStatus("Available")
	DiskStatusAttaching = DiskStatus("Attaching")
	DiskStatusDetaching = DiskStatus("Detaching")
	DiskStatusCreating  = DiskStatus("Creating")
	DiskStatusReIniting = DiskStatus("ReIniting")
)

type Disk ecs.DiskItemType

func (d *Disk) IsInUse() bool {
	return DiskStatus(d.Status) == DiskStatusInUse
}

func (d *Disk) IsAvailable() bool {
	return DiskStatus(d.Status) == DiskStatusAvailable
}

func (d *Disk) IsDetaching() bool {
	return DiskStatus(d.Status) == DiskStatusDetaching
}

func (d *Disk) LocalDevice() string {
	return strings.Replace(d.Device, DevicePrefixAPI, DevicePrefixLocal, 1)
}

func (d *Disk) SetLocalDevice(device string) {
	d.Device = strings.Replace(device, DevicePrefixLocal, DevicePrefixAPI, 1)
}
