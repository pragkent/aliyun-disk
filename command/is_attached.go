package command

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

type IsAttachedCommand struct {
	Meta
}

func (c *IsAttachedCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("is_attached")
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

	status := c.Meta.Driver.IsAttached(options, node)
	c.Meta.Ui.Output(jsonify(status))

	return 0
}

func (c *IsAttachedCommand) Synopsis() string {
	return "Check the volume is attached on the node"
}

func (c *IsAttachedCommand) Help() string {
	helpText := `
Usage: aliyun-disk isattached <json options> <node name>

Check the volume is attached on the node.
`
	return strings.TrimSpace(helpText)
}
