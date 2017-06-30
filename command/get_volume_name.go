package command

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

type GetVolumeNameCommand struct {
	Meta
}

func (c *GetVolumeNameCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("get_volume_name")
	if err := fs.Parse(args); err != nil {
		return cli.RunResultHelp
	}

	if fs.NArg() < 1 {
		return cli.RunResultHelp
	}

	options, err := volume.ParseOptions(fs.Arg(0))
	if err != nil {
		c.Meta.Ui.Output(jsonify(volume.NewDriverError(err)))
		return 1
	}

	status := c.Meta.Driver.GetVolumeName(options)
	c.Meta.Ui.Output(jsonify(status))

	return 1
}

func (c *GetVolumeNameCommand) Synopsis() string {
	return "Get volume name"
}

func (c *GetVolumeNameCommand) Help() string {
	helpText := `
Usage: aliyun-disk getvolumename <json options>

Get volume name.
`
	return strings.TrimSpace(helpText)
}
