package command

import (
	"strings"
)

type UnmountDeviceCommand struct {
	Meta
}

func (c *UnmountDeviceCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *UnmountDeviceCommand) Synopsis() string {
	return ""
}

func (c *UnmountDeviceCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
