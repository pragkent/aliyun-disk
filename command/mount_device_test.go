package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestMountDeviceCommand_implement(t *testing.T) {
	var _ cli.Command = &MountDeviceCommand{}
}
