package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAttachCommand_implement(t *testing.T) {
	var _ cli.Command = &AttachCommand{}
}
