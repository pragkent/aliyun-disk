package provider

import (
	"testing"

	"github.com/denverdino/aliyungo/ecs"
)

func TestIsDiskAttached(t *testing.T) {
	tests := []struct {
		instanceId     string
		diskInstanceId string
		status         DiskStatus
		want           bool
	}{
		{"12345", "12345", DiskStatusInUse, true},
		{"12345", "12345", DiskStatusDetaching, false},
		{"12345", "abcde", DiskStatusInUse, false},
	}

	for _, tt := range tests {
		instance := &Instance{InstanceId: tt.instanceId}
		disk := &Disk{InstanceId: tt.diskInstanceId, Status: ecs.DiskStatus(tt.status)}

		isAttached := instance.IsDiskAttached(disk)
		if isAttached != tt.want {
			t.Errorf("IsDiskAttached = %v; want: %v", isAttached, tt.want)
		}
	}
}
