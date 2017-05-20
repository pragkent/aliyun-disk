package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestUnmountCommand_implement(t *testing.T) {
	var _ cli.Command = &UnmountCommand{}
}
