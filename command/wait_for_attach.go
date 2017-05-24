package command

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

type WaitForAttachCommand struct {
	Meta
}

func (c *WaitForAttachCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("wait_for_attach")
	if err := fs.Parse(args); err != nil {
		return cli.RunResultHelp
	}

	if fs.NArg() < 2 {
		return cli.RunResultHelp
	}

	device := fs.Arg(0)
	options, err := volume.ParseOptions(fs.Arg(1))
	if err != nil {
		c.Meta.Ui.Output(jsonify(volume.NewDriverError(err)))
		return 1
	}

	status := c.Meta.Driver.WaitForAttach(device, options)
	c.Meta.Ui.Output(jsonify(status))
	return 0
}

func (c *WaitForAttachCommand) Synopsis() string {
	return "Wait for the volume to be attached on the remote node"
}

func (c *WaitForAttachCommand) Help() string {
	helpText := `
Usage: aliyun-disk waitforattach <device> <json options>

Wait for the volume to be attached on the remote node.
`
	return strings.TrimSpace(helpText)
}
