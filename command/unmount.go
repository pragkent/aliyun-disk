package command

import (
	"strings"
)

type UnmountCommand struct {
	Meta
}

func (c *UnmountCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *UnmountCommand) Synopsis() string {
	return ""
}

func (c *UnmountCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
