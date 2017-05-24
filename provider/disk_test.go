package provider

import "testing"

func TestLocalDevice(t *testing.T) {
	tests := []struct {
		apiDevice   string
		localDevice string
	}{
		{"/dev/xvda", "/dev/vda"},
		{"/dev/vda", "/dev/vda"},
	}

	for _, tt := range tests {
		d := Disk{Device: tt.apiDevice}
		if d.LocalDevice() != tt.localDevice {
			t.Errorf("LocalDevice error. Device = %s; LocalDevice = %s; want %s", tt.apiDevice, d.LocalDevice(), tt.localDevice)
		}
	}
}

func TestSetLocalDevice(t *testing.T) {
	tests := []struct {
		apiDevice   string
		localDevice string
	}{
		{"/dev/xvda", "/dev/vda"},
		{"/dev/xvda", "/dev/xvda"},
	}

	for _, tt := range tests {
		d := Disk{Device: ""}
		d.SetLocalDevice(tt.localDevice)

		if d.Device != tt.apiDevice {
			t.Errorf("SetLocalDevice error. LocalDevice = %s; Device = %s; want %s", tt.localDevice, d.Device, tt.apiDevice)
		}
	}
}
