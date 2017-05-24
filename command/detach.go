package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type DetachCommand struct {
	Meta
}

func (c *DetachCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("detach")
	if err := fs.Parse(args); err != nil {
		return cli.RunResultHelp
	}

	if fs.NArg() < 1 {
		return cli.RunResultHelp
	}

	device := fs.Arg(0)
	node, _ := getHostname()
	if fs.NArg() >= 2 {
		node = fs.Arg(1)
	}

	status := c.Meta.Driver.Detach(device, node)
	c.Meta.Ui.Output(jsonify(status))

	return 0
}

func (c *DetachCommand) Synopsis() string {
	return "Detach disk from ecs instance"
}

func (c *DetachCommand) Help() string {
	helpText := `
Usage: aliyun-disk detach <device> <node name>

Detach aliyun disk from ecs instance.
`
	return strings.TrimSpace(helpText)
}
