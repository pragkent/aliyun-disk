package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestMountCommand_implement(t *testing.T) {
	var _ cli.Command = &MountCommand{}
}
