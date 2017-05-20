package command

import (
	"strings"
)

type WaitForAttachCommand struct {
	Meta
}

func (c *WaitForAttachCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *WaitForAttachCommand) Synopsis() string {
	return ""
}

func (c *WaitForAttachCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
