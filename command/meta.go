package command

import (
	"errors"

	"github.com/mitchellh/cli"
)

// Meta contain the meta-option that nearly all subcommand inherited.
type Meta struct {
	Ui cli.Ui

	AccessKey string
	SecretKey string
}

func (m *Meta) validateCredentials() error {
	if m.AccessKey == "" {
		return errors.New("AccessKey is missing")
	}

	if m.SecretKey == "" {
		return errors.New("SecretKey is missing")
	}

	return nil
}
