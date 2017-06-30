package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/command"
	_ "github.com/pragkent/aliyun-disk/logs"
	"github.com/pragkent/aliyun-disk/volume"
)

func main() {
	args := os.Args[1:]
	meta := newMeta()

	cli := &cli.CLI{
		Args:       args,
		Commands:   Commands(meta),
		Version:    Version,
		HelpFunc:   cli.BasicHelpFunc(Name),
		HelpWriter: os.Stdout,
	}

	exitStatus, err := cli.Run()
	if err != nil {
		meta.Ui.Error(err.Error())
	}
	os.Exit(exitStatus)
}

func newMeta() *command.Meta {
	ui := &cli.BasicUi{
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
		Reader:      os.Stdin,
	}

	driver := volume.NewDriver(getDriverConfig())
	return &command.Meta{
		Ui:     ui,
		Driver: driver,
	}
}

func getDriverConfig() *volume.DriverConfig {
	return &volume.DriverConfig{
		Region:    os.Getenv("ALIYUN_REGION"),
		AccessKey: os.Getenv("ALIYUN_ACCESS_KEY"),
		SecretKey: os.Getenv("ALIYUN_SECRET_KEY"),
		Cluster:   os.Getenv("ALIYUN_CLUSTER"),
	}
}
