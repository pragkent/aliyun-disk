package command

import (
	"strings"
)

type GetVolumeNameCommand struct {
	Meta
}

func (c *GetVolumeNameCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GetVolumeNameCommand) Synopsis() string {
	return ""
}

func (c *GetVolumeNameCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
