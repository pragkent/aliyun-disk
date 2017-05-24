package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type UnmountCommand struct {
	Meta
}

func (c *UnmountCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("unmount")
	if err := fs.Parse(args); err != nil {
		return cli.RunResultHelp
	}

	if fs.NArg() < 1 {
		return cli.RunResultHelp
	}

	dir := fs.Arg(0)

	status := c.Meta.Driver.Unmount(dir)
	c.Meta.Ui.Output(jsonify(status))

	return 1
}

func (c *UnmountCommand) Synopsis() string {
	return "Unmount the volume"
}

func (c *UnmountCommand) Help() string {
	helpText := `
Usage: aliyun-disk unmount <mount dir>

Unmount the volume.
`
	return strings.TrimSpace(helpText)
}
