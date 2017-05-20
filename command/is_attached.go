package command

import (
	"strings"
)

type IsAttachedCommand struct {
	Meta
}

func (c *IsAttachedCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *IsAttachedCommand) Synopsis() string {
	return ""
}

func (c *IsAttachedCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
