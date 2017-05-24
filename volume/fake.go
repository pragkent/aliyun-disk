package volume

type FakeDriver struct{}

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
		VolumeName: "xyz",
	}
}

func (d *FakeDriver) Attach(options Options, node string) *DriverStatus {
	return &DriverStatus{
		Status:     StatusSuccess,
		DevicePath: "/dev/vdb",
	}
}

func (d *FakeDriver) Detach(device string, node string) *DriverStatus {
	return &DriverStatus{
		Status: StatusSuccess,
	}
}

func (d *FakeDriver) WaitForAttach(device string, options Options) *DriverStatus {
	return &DriverStatus{
		Status:     StatusSuccess,
		DevicePath: "/dev/vdb",
	}
}

func (d *FakeDriver) IsAttached(options Options, node string) *DriverStatus {
	return &DriverStatus{
		Status:   StatusSuccess,
		Attached: true,
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
