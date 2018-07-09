package provider

import "github.com/denverdino/aliyungo/ecs"

type Instance ecs.InstanceAttributesType

func (i *Instance) IsDiskAttached(disk *Disk) bool {
	return i.InstanceId == disk.InstanceId && disk.IsInUse()
}
