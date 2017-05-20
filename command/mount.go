package command

import (
	"strings"
)

type MountCommand struct {
	Meta
}

func (c *MountCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *MountCommand) Synopsis() string {
	return ""
}

func (c *MountCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
