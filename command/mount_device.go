package command

import (
	"strings"
)

type MountDeviceCommand struct {
	Meta
}

func (c *MountDeviceCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *MountDeviceCommand) Synopsis() string {
	return ""
}

func (c *MountDeviceCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
