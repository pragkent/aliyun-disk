package volume

import (
	"fmt"
	"log"
	"os"

	"k8s.io/kubernetes/pkg/util/exec"
	"k8s.io/kubernetes/pkg/util/mount"

	"github.com/pragkent/aliyun-disk/provider"
)

type AliyunDriver struct {
	aliyun provider.Provider
}

func NewDriver(accessKey string, secretKey string, region string) Driver {
	if accessKey == "" || secretKey == "" || region == "" {
		return &AliyunDriver{}
	}

	return &AliyunDriver{
		aliyun: provider.New(accessKey, secretKey, region),
	}
}

func (d *AliyunDriver) isAliyunAPIAvailable() bool {
	return d.aliyun != nil
}

func (d *AliyunDriver) Init() *DriverStatus {
	log.Printf("Init invoked.")
	return &DriverStatus{
		Status: StatusSuccess,
	}
}

func (d *AliyunDriver) GetVolumeName(options Options) *DriverStatus {
	log.Printf("GetVolumeName invoked. Options: %v", options)

	diskId := options.DiskId()
	return &DriverStatus{
		Status:     StatusSuccess,
		VolumeName: diskId,
	}
}

func (d *AliyunDriver) Attach(options Options, node string) *DriverStatus {
	log.Printf("Attach invoked. Options: %v Node: %s", options, node)

	if !d.isAliyunAPIAvailable() {
		return NewDriverError(errAliyunAPIUnavailable)
	}

	instance, err := d.aliyun.GetInstanceByHostname(node)
	if err != nil {
		log.Printf("GetInstanceByHostname failed. %v. Hostname: %s", err, node)
		return NewDriverError(fmt.Errorf("could not find instance %q: %v", node, err))
	}

	diskId := options.DiskId()
	disk, err := d.aliyun.GetDiskById(diskId)
	if err != nil {
		log.Printf("GetDiskById failed. %v. DiskId: %s", err, diskId)
		return NewDriverError(fmt.Errorf("could not find disk %q: %v", diskId, err))
	}

	if instance.IsDiskAttached(disk) {
		log.Printf("Disk already attached. DiskId: %s", diskId)
		return &DriverStatus{
			Status:     StatusSuccess,
			DevicePath: disk.LocalDevice(),
		}
	}

	if disk.IsInUse() {
		log.Printf("Try to detach InUse disk from instance. DiskId: %s InstanceId: %s",
			diskId, disk.InstanceId)
		if err := d.detachDisk(disk); err != nil {
			log.Printf("DetachDisk failed. %v. DiskId: %s", err, diskId)
			return NewDriverError(err)
		}
	}

	if err := d.attachDisk(instance, disk); err != nil {
		log.Printf("AttachDisk failed. %v. DiskId: %s InstanceId: %s", err, diskId, instance.InstanceId)
		return NewDriverError(err)
	}

	disk, err = d.aliyun.GetDiskById(disk.DiskId)
	if err != nil {
		log.Printf("GetDiskById failed. %v. DiskId: %s", err, disk.DiskId)
		return NewDriverError(fmt.Errorf("could not get disk %q: %v", disk.DiskId, err))
	}

	log.Printf("Attach finished. DiskId: %s InstanceId: %s Node: %s",
		disk.DiskId, instance.InstanceId, node)

	return &DriverStatus{
		Status:     StatusSuccess,
		DevicePath: disk.LocalDevice(),
	}
}

func (d *AliyunDriver) detachDisk(disk *provider.Disk) error {
	err := d.aliyun.DetachDisk(disk.InstanceId, disk.DiskId)
	if err != nil {
		return fmt.Errorf("Unable to detach disk %s: %v", disk.DiskId, err)
	}

	return d.aliyun.WaitForDisk(disk.DiskId, provider.DiskStatusAvailable)
}

func (d *AliyunDriver) attachDisk(instance *provider.Instance, disk *provider.Disk) error {
	if err := d.aliyun.AttachDisk(instance.InstanceId, disk.DiskId); err != nil {
		return fmt.Errorf("Unable to attach disk %s: %v", disk.DiskId, err)
	}

	return d.aliyun.WaitForDisk(disk.DiskId, provider.DiskStatusInUse)
}

func (d *AliyunDriver) Detach(device string, node string) *DriverStatus {
	log.Printf("Detach invoked. Device: %s Node: %s", device, node)

	if !d.isAliyunAPIAvailable() {
		return NewDriverError(errAliyunAPIUnavailable)
	}

	instance, err := d.aliyun.GetInstanceByHostname(node)
	if err != nil {
		log.Printf("GetInstanceByHostname failed. %v. Hostname: %s", err, node)
		return NewDriverError(fmt.Errorf("could not find instance %q: %v", node, err))
	}

	disk, err := d.aliyun.GetDiskById(device)
	if err != nil {
		log.Printf("GetDiskById failed. %v. InstanceId: %s Device: %s", err,
			instance.InstanceId, device)
		return NewDriverError(fmt.Errorf("could not find disk %q: %v", device, err))
	}

	if disk.IsDetaching() || disk.IsAvailable() {
		return &DriverStatus{
			Status: StatusSuccess,
		}
	}

	if !instance.IsDiskAttached(disk) {
		return &DriverStatus{
			Status: StatusSuccess,
		}
	}

	if err := d.aliyun.DetachDisk(instance.InstanceId, disk.DiskId); err != nil {
		log.Printf("DetachDisk failed. %v. InstanceId: %s DiskId: %s", err,
			instance.InstanceId, disk.DiskId)
		return NewDriverError(fmt.Errorf("could not detach disk %q: %v", disk.DiskId, err))
	}

	log.Printf("Detach finished. InstanceID: %s DiskId: %s Node: %s Device: %s",
		instance.InstanceId, disk.DiskId, node, device)
	return &DriverStatus{
		Status: StatusSuccess,
	}
}

func (d *AliyunDriver) WaitForAttach(device string, options Options) *DriverStatus {
	log.Printf("WaitForAttach invoked. Device: %s Options: %v", device, options)

	if !d.isAliyunAPIAvailable() {
		return NewDriverNotSupported(errAliyunAPIUnavailable)
	}

	diskId := options.DiskId()
	disk, err := d.aliyun.GetDiskById(diskId)
	if err != nil {
		return NewDriverError(fmt.Errorf("could not find disk %q: %v", diskId, err))
	}

	if err := d.aliyun.WaitForDisk(disk.DiskId, provider.DiskStatusInUse); err != nil {
		return NewDriverError(err)
	}

	disk, err = d.aliyun.GetDiskById(diskId)
	if err != nil {
		return NewDriverError(fmt.Errorf("could not find disk %q: %v", diskId, err))
	}

	log.Printf("WaitForAttach finished. Device: %s DiskId: %s", device, diskId)
	return &DriverStatus{
		Status:     StatusSuccess,
		DevicePath: disk.LocalDevice(),
	}
}

func (d *AliyunDriver) IsAttached(options Options, node string) *DriverStatus {
	log.Printf("IsAttached invoked. Options: %v Node: %s", options, node)

	if !d.isAliyunAPIAvailable() {
		return NewDriverNotSupported(errAliyunAPIUnavailable)
	}

	instance, err := d.aliyun.GetInstanceByHostname(node)
	if err != nil {
		return NewDriverError(fmt.Errorf("could not find instance %q: %v", node, err))
	}

	diskId := options.DiskId()
	disk, err := d.aliyun.GetDiskById(diskId)
	if err != nil {
		return NewDriverError(fmt.Errorf("could not find disk %q: %v", diskId, err))
	}

	return &DriverStatus{
		Status:   StatusSuccess,
		Attached: instance.IsDiskAttached(disk) && disk.IsInUse(),
	}
}

func (d *AliyunDriver) MountDevice(dir string, device string, options Options) *DriverStatus {
	log.Printf("MountDevice invoked. Dir: %s Device: %s Options: %v", dir, device, options)

	fsType, _ := options[optionFSType].(string)
	rw, _ := options[optionReadWrite].(string)

	flags := []string{}
	if rw != "" {
		flags = append(flags, rw)
	}

	notMnt, err := mount.IsNotMountPoint(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0750); err != nil {
				return NewDriverError(err)
			}
			notMnt = true
		} else {
			return NewDriverError(err)
		}
	}

	if notMnt {
		mounter := &mount.SafeFormatAndMount{Interface: mount.New(""), Runner: exec.New()}
		if err := mounter.FormatAndMount(device, dir, fsType, flags); err != nil {
			os.Remove(dir)
			return NewDriverError(fmt.Errorf("FormatAndMount error: %v", err))
		}
	}

	log.Printf("MountDevice finished. Dir: %s Device: %s fsType: %s Flags: %s",
		dir, device, fsType, flags)
	return &DriverStatus{
		Status: StatusSuccess,
	}
}

func (d *AliyunDriver) UnmountDevice(dir string) *DriverStatus {
	log.Printf("UnmountDevice invoked. Dir: %s", dir)

	mounter := &mount.SafeFormatAndMount{Interface: mount.New(""), Runner: exec.New()}
	if err := UnmountPath(dir, mounter); err != nil {
		log.Printf("UnmountPath failed. %s. Dir: %s", err, dir)
		return NewDriverError(fmt.Errorf("UnmountPath failed: %v", err))
	}

	log.Printf("UnmountDevice finished. Dir: %s", dir)
	return &DriverStatus{
		Status: StatusSuccess,
	}
}

func (d *AliyunDriver) Mount(dir string, options Options) *DriverStatus {
	log.Printf("Mount invoked. Dir: %s Options: %v", dir, options)
	return NewDriverNotSupported(errCommandNotSupported)
}

func (d *AliyunDriver) Unmount(dir string) *DriverStatus {
	log.Printf("Unmount invoked. Dir: %s", dir)
	return NewDriverNotSupported(errCommandNotSupported)
}
