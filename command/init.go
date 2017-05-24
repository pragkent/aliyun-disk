package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type InitCommand struct {
	Meta
}

func (c *InitCommand) Run(args []string) int {
	fs := c.Meta.FlagSet("init")
	if err := fs.Parse(args); err != nil {
		return cli.RunResultHelp
	}

	status := c.Meta.Driver.Init()
	c.Meta.Ui.Output(jsonify(status))

	return 0
}

func (c *InitCommand) Synopsis() string {
	return "Initialize aliyun disk driver"
}

func (c *InitCommand) Help() string {
	helpText := `
Usage: aliyun-disk init

Initialize aliyun disk flexvolume driver
`
	return strings.TrimSpace(helpText)
}
