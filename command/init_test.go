package command

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

func TestInitCommand_Run(t *testing.T) {
	var bu bytes.Buffer
	ui := &cli.BasicUi{
		Writer: &bu,
	}

	meta := &Meta{
		Ui:     ui,
		Driver: volume.NewFakeDriver(),
	}

	cmd := &InitCommand{*meta}
	if status := cmd.Run(nil); status != 0 {
		t.Errorf("cmd.Run() == %d; want 0", status)
	}

	var ds volume.DriverStatus
	if err := json.Unmarshal(bu.Bytes(), &ds); err != nil {
		t.Errorf("json.Unmarshal error. %v", err)
	}

	if ds.Status != volume.StatusSuccess {
		t.Errorf("VolumeStatus = %s; want %s", ds.Status, volume.StatusSuccess)
	}
}
