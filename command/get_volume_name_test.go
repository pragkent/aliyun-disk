package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGetVolumeNameCommand_implement(t *testing.T) {
	var _ cli.Command = &GetVolumeNameCommand{}
}
