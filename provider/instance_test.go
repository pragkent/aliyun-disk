package provider

import "testing"

func TestIsDiskAttached(t *testing.T) {
	tests := []struct {
		instanceId     string
		diskInstanceId string
		want           bool
	}{
		{"12345", "12345", true},
		{"12345", "abcde", false},
	}

	for _, tt := range tests {
		instance := &Instance{InstanceId: tt.instanceId}
		disk := &Disk{InstanceId: tt.diskInstanceId}

		isAttached := instance.IsDiskAttached(disk)
		if isAttached != tt.want {
			t.Errorf("IsDiskAttached = %v; want: %v", isAttached, tt.want)
		}
	}
}
