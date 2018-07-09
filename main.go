package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	cfg := loadDriverConfig()
	if cfg != nil {
		return cfg
	}

	return &volume.DriverConfig{
		Region:    os.Getenv("ALIYUN_REGION"),
		AccessKey: os.Getenv("ALIYUN_ACCESS_KEY"),
		SecretKey: os.Getenv("ALIYUN_SECRET_KEY"),
		Cluster:   os.Getenv("ALIYUN_CLUSTER"),
	}
}

func loadDriverConfig() *volume.DriverConfig {
	const path = "/etc/kubernetes/cloud/cloud.json"
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("ioutil.ReadFile error: %v", err)
		return nil
	}

	var cfg volume.DriverConfig
	if err = json.Unmarshal(raw, &cfg); err != nil {
		log.Printf("json.Unmarshal error: %v", err)
		return nil
	}

	return &cfg
}
