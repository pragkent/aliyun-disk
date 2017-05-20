package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestIsAttachedCommand_implement(t *testing.T) {
	var _ cli.Command = &IsAttachedCommand{}
}
