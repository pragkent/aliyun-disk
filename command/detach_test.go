package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDetachCommand_implement(t *testing.T) {
	var _ cli.Command = &DetachCommand{}
}
