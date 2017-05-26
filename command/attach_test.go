package command

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

func TestAttachCommand_Run_ArgsError(t *testing.T) {
	var bu bytes.Buffer
	ui := &cli.BasicUi{
		Writer: &bu,
	}

	meta := &Meta{
		Ui:     ui,
		Driver: volume.NewFakeDriver(),
	}

	tests := []struct {
		args   []string
		result int
	}{
		{nil, cli.RunResultHelp},
		{[]string{"-x"}, cli.RunResultHelp},
	}

	for _, tt := range tests {
		cmd := &AttachCommand{*meta}
		if result := cmd.Run(tt.args); result != result {
			t.Errorf("cmd.Run() == %d; want %d", result, result)
		}
	}
}

func TestAttachCommand_Run(t *testing.T) {
	var bu bytes.Buffer
	ui := &cli.BasicUi{
		Writer: &bu,
	}

	meta := &Meta{
		Ui:     ui,
		Driver: volume.NewFakeDriver(),
	}

	cmd := &AttachCommand{*meta}

	tests := []struct {
		args   []string
		result int
		status volume.DriverStatus
	}{
		{
			[]string{"nojson", "node0"},
			1,
			volume.DriverStatus{
				Status:  volume.StatusFailure,
				Message: "invalid volume options",
			},
		},
		{
			[]string{`{"abc": "efg"}`, "node0"},
			1,
			volume.DriverStatus{
				Status:  volume.StatusFailure,
				Message: "option diskId is required",
			},
		},
		{
			[]string{`{"diskId": "12345"}`, "node0"},
			0,
			volume.DriverStatus{
				Status:     volume.StatusSuccess,
				DevicePath: "/dev/vdb",
			},
		},
		{
			[]string{`{"diskId": "12345"}`, "node1"},
			0,
			volume.DriverStatus{
				Status:     volume.StatusSuccess,
				DevicePath: "/dev/xyz",
			},
		},
		{
			[]string{`{"diskId": "12345"}`, "node2"},
			0,
			volume.DriverStatus{
				Status:  volume.StatusFailure,
				Message: "unknown node",
			},
		},
	}

	for _, tt := range tests {
		bu.Reset()

		if result := cmd.Run(tt.args); result != tt.result {
			t.Errorf("cmd.Run() == %d; want %d", result, tt.result)
		}

		var ds volume.DriverStatus
		if err := json.Unmarshal(bu.Bytes(), &ds); err != nil {
			t.Errorf("json.Unmarshal error. %v", err)
		}

		if !reflect.DeepEqual(ds, tt.status) {
			t.Errorf("Status = %#v; want %#v", ds, tt.status)
		}
	}
}
