package provider

import "time"

type Provider interface {
	GetInstanceByHostname(hostname string) (*Instance, error)
	GetDiskById(diskId string) (*Disk, error)
	GetDiskByTags(tags map[string]string) (*Disk, error)
	AttachDisk(instanceId string, diskId string) error
	DetachDisk(isntanceId string, diskId string) error
	WaitForDisk(diskId string, status DiskStatus, timeout time.Duration) error
	AddTags(args *AddTagsArgs) error
}
