package command

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

type MountDeviceCommand struct {
	Meta
}

func (c *MountDeviceCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("mount_device")
	if err := fs.Parse(args); err != nil {
		return cli.RunResultHelp
	}

	if fs.NArg() < 3 {
		return cli.RunResultHelp
	}

	dir := fs.Arg(0)
	device := fs.Arg(1)
	options, err := volume.ParseOptions(fs.Arg(2))
	if err != nil {
		c.Meta.Ui.Output(jsonify(volume.NewDriverError(err)))
		return 1
	}

	status := c.Meta.Driver.MountDevice(dir, device, options)
	c.Meta.Ui.Output(jsonify(status))

	return 0
}

func (c *MountDeviceCommand) Synopsis() string {
	return "Mount device to a global path"
}

func (c *MountDeviceCommand) Help() string {
	helpText := `
Usage: aliyun-disk mountdevice <mount dir> <mount device> <json options>

Mount device to a global path.
`
	return strings.TrimSpace(helpText)
}
