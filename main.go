package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/command"
)

func main() {
	args := fixArgs(os.Args)
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

func fixArgs(origin []string) []string {
	args := origin[1:]

	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	return args
}

func newMeta() *command.Meta {
	return &command.Meta{
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
			Reader:      os.Stdin,
		},
		AccessKey: os.Getenv("ALIYUN_ACCESS_KEY"),
		SecretKey: os.Getenv("ALIYUN_SECRET_KEY"),
	}
}
