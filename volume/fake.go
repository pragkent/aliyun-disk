package volume

type FakeDriver struct{}

var nodeDeviceMap = map[string]string{
	"node0": "/dev/vdb",
	"node1": "/dev/xyz",
}

func NewFakeDriver() Driver {
	return &FakeDriver{}
}

func (d *FakeDriver) Init() *DriverStatus {
	return &DriverStatus{
		Status: StatusSuccess,
	}
}

func (d *FakeDriver) GetVolumeName(options Options) *DriverStatus {
	return &DriverStatus{
		Status:     StatusSuccess,
		VolumeName: options.DiskId(),
	}
}

func (d *FakeDriver) Attach(options Options, node string) *DriverStatus {
	if dp, ok := nodeDeviceMap[node]; ok {
		return &DriverStatus{
			Status:     StatusSuccess,
			DevicePath: dp,
		}
	}

	return &DriverStatus{
		Status:  StatusFailure,
		Message: "unknown node",
	}
}

func (d *FakeDriver) Detach(device string, node string) *DriverStatus {
	return &DriverStatus{
		Status: StatusSuccess,
	}
}

func (d *FakeDriver) WaitForAttach(device string, options Options) *DriverStatus {
	for _, d := range nodeDeviceMap {
		if d == device {
			return &DriverStatus{
				Status:     StatusSuccess,
				DevicePath: d,
			}
		}
	}

	return &DriverStatus{
		Status:  StatusFailure,
		Message: "unknown device",
	}
}

func (d *FakeDriver) IsAttached(options Options, node string) *DriverStatus {
	if _, ok := nodeDeviceMap[node]; ok {
		return &DriverStatus{
			Status:   StatusSuccess,
			Attached: true,
		}
	}

	return &DriverStatus{
		Status:   StatusSuccess,
		Attached: false,
	}
}

func (d *FakeDriver) MountDevice(dir string, device string, options Options) *DriverStatus {
	return &DriverStatus{
		Status: StatusSuccess,
	}
}

func (d *FakeDriver) UnmountDevice(dir string) *DriverStatus {
	return &DriverStatus{
		Status: StatusSuccess,
	}
}

func (d *FakeDriver) Mount(dir string, options Options) *DriverStatus {
	return NewDriverNotSupported(errCommandNotSupported)
}

func (d *FakeDriver) Unmount(dir string) *DriverStatus {
	return NewDriverNotSupported(errCommandNotSupported)
}
