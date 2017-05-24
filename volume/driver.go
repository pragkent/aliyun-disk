package volume

type Driver interface {
	Init() *DriverStatus
	GetVolumeName(options Options) *DriverStatus
	Attach(options Options, node string) *DriverStatus
	Detach(device string, node string) *DriverStatus
	WaitForAttach(device string, options Options) *DriverStatus
	IsAttached(options Options, node string) *DriverStatus
	MountDevice(dir string, device string, options Options) *DriverStatus
	UnmountDevice(dir string) *DriverStatus
	Mount(dir string, options Options) *DriverStatus
	Unmount(dir string) *DriverStatus
}
