package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type UnmountDeviceCommand struct {
	Meta
}

func (c *UnmountDeviceCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("unmount_device")
	if err := fs.Parse(args); err != nil {
		return cli.RunResultHelp
	}

	if fs.NArg() < 1 {
		return cli.RunResultHelp
	}

	device := fs.Arg(0)

	status := c.Meta.Driver.UnmountDevice(device)
	c.Meta.Ui.Output(jsonify(status))

	return 0
}

func (c *UnmountDeviceCommand) Synopsis() string {
	return "Unmount the global mount for the device."
}

func (c *UnmountDeviceCommand) Help() string {
	helpText := `
Usage: aliyun-disk unmountdevice <mount device>

Unmount the global mount for the device.
`
	return strings.TrimSpace(helpText)
}
