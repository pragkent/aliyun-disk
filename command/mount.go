package command

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

type MountCommand struct {
	Meta
}

func (c *MountCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("mount")
	if err := fs.Parse(args); err != nil {
		return cli.RunResultHelp
	}

	if fs.NArg() < 2 {
		return cli.RunResultHelp
	}

	dir := fs.Arg(0)
	options, err := volume.ParseOptions(fs.Arg(1))
	if err != nil {
		c.Meta.Ui.Output(jsonify(volume.NewDriverError(err)))
		return 1
	}

	status := c.Meta.Driver.Mount(dir, options)
	c.Meta.Ui.Output(jsonify(status))

	return 1
}

func (c *MountCommand) Synopsis() string {
	return "Mount the volume at the mount dir"
}

func (c *MountCommand) Help() string {
	helpText := `
Usage: aliyun-disk mount <mount dir> <json options>

Mount the volume at the mount dir.
`
	return strings.TrimSpace(helpText)
}
