package command

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

func TestMountCommand_Run_ArgsError(t *testing.T) {
	ui := &cli.BasicUi{}

	meta := &Meta{
		Ui:     ui,
		Driver: volume.NewFakeDriver(),
	}

	cmd := &MountCommand{*meta}

	tests := []struct {
		args   []string
		result int
	}{
		{[]string{}, cli.RunResultHelp},
		{[]string{"dir"}, cli.RunResultHelp},
	}

	for _, tt := range tests {
		if result := cmd.Run(tt.args); result != tt.result {
			t.Errorf("cmd.Run() == %d; want %d", result, tt.result)
		}
	}

}
func TestMountCommand_Run(t *testing.T) {
	var bu bytes.Buffer
	ui := &cli.BasicUi{
		Writer: &bu,
	}

	meta := &Meta{
		Ui:     ui,
		Driver: volume.NewFakeDriver(),
	}

	cmd := &MountCommand{*meta}

	tests := []struct {
		args   []string
		result int
		status volume.DriverStatus
	}{
		{
			[]string{"/mnt/xx", "no_json"},
			1,
			volume.DriverStatus{
				Status:  volume.StatusFailure,
				Message: "invalid volume options",
			},
		},
		{
			[]string{"/mnt/xx", `{"abc": "efg"}`},
			1,
			volume.DriverStatus{
				Status:  volume.StatusFailure,
				Message: "option diskId is required",
			},
		},
		{
			[]string{"/mnt/xx", `{"diskId": "12345"}`},
			1,
			volume.DriverStatus{
				Status:  volume.StatusNotSupported,
				Message: "command not supported",
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
