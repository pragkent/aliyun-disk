package command

import (
	"strings"
)

type AttachCommand struct {
	Meta
}

func (c *AttachCommand) Run(args []string) int {
	if err := c.Meta.validateCredentials(); err != nil {
		c.Meta.Ui.Error(err.Error())
		return 1
	}

	return 0
}

func (c *AttachCommand) Synopsis() string {
	return "Attach disk to ecs instance"
}

func (c *AttachCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
