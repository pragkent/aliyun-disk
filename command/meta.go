package command

import (
	"bufio"
	"flag"
	"io"

	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/volume"
)

// Meta contain the meta-option that nearly all subcommand inherited.
type Meta struct {
	Ui cli.Ui

	Driver volume.Driver
}

func (m *Meta) FlagSet(name string) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)

	// Create an io.Writer that writes to our Ui properly for errors.
	// This is kind of a hack, but it does the job. Basically: create
	// a pipe, use a scanner to break it into lines, and output each line
	// to the UI. Do this forever.
	errR, errW := io.Pipe()
	errScanner := bufio.NewScanner(errR)
	go func() {
		for errScanner.Scan() {
			m.Ui.Error(errScanner.Text())
		}
	}()

	fs.SetOutput(errW)

	fs.Usage = func() {}

	return fs
}
