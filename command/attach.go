package command

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

type AttachCommand struct {
	Meta
}

func (c *AttachCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("attach")
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

	node, _ := getHostname()
	if fs.NArg() >= 2 {
		node = fs.Arg(1)
	}

	status := c.Meta.Driver.Attach(options, node)
	c.Meta.Ui.Output(jsonify(status))

	return 0
}

func (c *AttachCommand) Synopsis() string {
	return "Attach disk to ecs instance"
}

func (c *AttachCommand) Help() string {
	helpText := `
Usage: aliyun-disk attach <json options> <node name>

Attach aliyun disk to ecs instance.
`
	return strings.TrimSpace(helpText)
}
