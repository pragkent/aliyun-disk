package command

import (
	"strings"
)

type DetachCommand struct {
	Meta
}

func (c *DetachCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DetachCommand) Synopsis() string {
	return ""
}

func (c *DetachCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
