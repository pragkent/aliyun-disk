package provider

type Provider interface {
	GetInstanceByHostname(hostname string) (*Instance, error)
	GetDiskById(diskId string) (*Disk, error)
	AttachDisk(instanceId string, diskId string) error
	DetachDisk(isntanceId string, diskId string) error
	WaitForDisk(diskId string, status DiskStatus) error
}
