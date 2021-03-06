package provider

import (
	"errors"
	"time"

	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
)

const DefaultTimeout = time.Minute
const DefaultInterval = time.Second

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
	pg := &common.Pagination{
		PageNumber: 0,
		PageSize:   50,
	}

	var matched []ecs.InstanceAttributesType

	for pg != nil {
		args := &ecs.DescribeInstancesArgs{
			RegionId:   p.region,
			Pagination: *pg,
		}

		instances, pgr, err := p.c.DescribeInstances(args)
		if err != nil {
			return nil, err
		}

		for _, i := range instances {
			if i.HostName == hostname {
				matched = append(matched, i)
			}
		}

		pg = pgr.NextPage()
	}

	if len(matched) == 0 {
		return nil, ErrInstanceNotExist
	}

	if len(matched) > 1 {
		return nil, ErrHostnameDuplicated
	}

	instance := Instance(matched[0])
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
		return nil, ErrDiskNotExist
	}

	disk := Disk(ds[0])
	return &disk, nil
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

func (p *Aliyun) WaitForDisk(diskId string, status DiskStatus, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	args := ecs.DescribeDisksArgs{
		RegionId: p.region,
		DiskIds:  []string{diskId},
	}

	for {
		disks, _, err := p.c.DescribeDisks(&args)
		if err != nil {
			return err
		}

		if disks == nil || len(disks) == 0 {
			return errors.New("provider: disk not found")
		}

		if disks[0].Status == ecs.DiskStatus(status) {
			break
		}

		timeout -= DefaultInterval
		if timeout <= 0 {
			return errors.New("provider: wait for disk timeout")
		}

		time.Sleep(DefaultInterval)
	}

	return nil
}

func (p *Aliyun) AddTags(args *AddTagsArgs) error {
	ecsArgs := &ecs.AddTagsArgs{
		RegionId:     p.region,
		ResourceId:   args.ResourceId,
		ResourceType: ecs.TagResourceType(args.ResourceType),
		Tag:          args.Tag,
	}

	return p.c.AddTags(ecsArgs)
}

func (p *Aliyun) GetDiskByTags(tags map[string]string) (*Disk, error) {
	args := &ecs.DescribeDisksArgs{
		RegionId: p.region,
		Tag:      tags,
	}

	ds, _, err := p.c.DescribeDisks(args)
	if err != nil {
		return nil, err
	}

	if len(ds) == 0 {
		return nil, ErrDiskNotExist
	}

	if len(ds) > 1 {
		return nil, ErrDiskTagDuplicated
	}

	disk := Disk(ds[0])
	return &disk, nil
}
