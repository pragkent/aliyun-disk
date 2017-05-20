package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestWaitForAttachCommand_implement(t *testing.T) {
	var _ cli.Command = &WaitForAttachCommand{}
}
