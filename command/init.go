package command

import (
	"encoding/json"
	"strings"

	"github.com/pragkent/aliyun-disk/pkg/volume/flexvolume"
)

type InitCommand struct {
	Meta
}

func (c *InitCommand) Run(args []string) int {
	status := &flexvolume.DriverStatus{
		Status: flexvolume.StatusSuccess,
	}

	b, err := json.Marshal(status)
	if err != nil {
		c.Meta.Ui.Error(err.Error())
		return 1
	}

	c.Meta.Ui.Output(string(b))
	return 0
}

func (c *InitCommand) Synopsis() string {
	return "Initialize driver"
}

func (c *InitCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
