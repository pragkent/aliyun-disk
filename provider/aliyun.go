package provider

import (
	"fmt"

	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
)

type Aliyun struct {
	region common.Region
	c      *ecs.Client
}

func New(accessKey, secretKey, region string) Provider {
	return &Aliyun{
		region: common.Region(region),
		c:      ecs.NewECSClient(accessKey, secretKey, common.Region(region)),
	}
}

func (p *Aliyun) GetInstanceByHostname(hostname string) (*Instance, error) {
	tags := map[string]string{
		"hostname": hostname,
	}

	args := &ecs.DescribeInstancesArgs{
		RegionId: p.region,
		Tag:      tags,
	}

	ins, _, err := p.c.DescribeInstances(args)
	if err != nil {
		return nil, err
	}

	if len(ins) == 0 {
		return nil, errInstanceNotExist
	}

	if len(ins) > 1 {
		return nil, errHostnameDuplicated
	}

	instance := Instance(ins[0])
	return &instance, nil
}

func (p *Aliyun) GetDiskById(diskId string) (*Disk, error) {
	args := &ecs.DescribeDisksArgs{
		RegionId: p.region,
		DiskIds:  []string{diskId},
	}

	ds, _, err := p.c.DescribeDisks(args)
	if err != nil {
		return nil, err
	}

	if len(ds) == 0 {
		return nil, errDiskNotExist
	}

	disk := Disk(ds[0])
	return &disk, nil
}

func (p *Aliyun) GetDiskByDevice(instanceId string, device string) (*Disk, error) {
	pg := &common.Pagination{
		PageNumber: 0,
		PageSize:   50,
	}

	for pg != nil {
		args := &ecs.DescribeDisksArgs{
			RegionId:   common.Region(p.region),
			InstanceId: instanceId,
			Pagination: *pg,
		}

		disks, pgr, err := p.c.DescribeDisks(args)
		if err != nil {
			return nil, err
		}

		for _, d := range disks {
			disk := Disk(d)
			if disk.LocalDevice() == device {
				return &disk, nil
			}
		}

		pg = pgr.NextPage()
	}

	return nil, fmt.Errorf("Device %q is not found on instance %q", device, instanceId)
}

func (p *Aliyun) AttachDisk(instanceId string, diskId string) error {
	args := &ecs.AttachDiskArgs{
		InstanceId: instanceId,
		DiskId:     diskId,
	}

	return p.c.AttachDisk(args)
}

func (p *Aliyun) DetachDisk(instanceId string, diskId string) error {
	return p.c.DetachDisk(instanceId, diskId)
}

func (p *Aliyun) WaitForDisk(diskId string, status DiskStatus) error {
	return p.c.WaitForDisk(p.region, diskId, ecs.DiskStatus(status), 0)
}
