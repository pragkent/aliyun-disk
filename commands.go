package main

import (
	"github.com/mitchellh/cli"
	"github.com/pragkent/aliyun-disk/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &command.InitCommand{
				Meta: *meta,
			}, nil
		},
		"attach": func() (cli.Command, error) {
			return &command.AttachCommand{
				Meta: *meta,
			}, nil
		},
		"isattached": func() (cli.Command, error) {
			return &command.IsAttachedCommand{
				Meta: *meta,
			}, nil
		},
		"detach": func() (cli.Command, error) {
			return &command.DetachCommand{
				Meta: *meta,
			}, nil
		},
		"mountdevice": func() (cli.Command, error) {
			return &command.MountDeviceCommand{
				Meta: *meta,
			}, nil
		},
		"unmountdevice": func() (cli.Command, error) {
			return &command.UnmountDeviceCommand{
				Meta: *meta,
			}, nil
		},
		"mount": func() (cli.Command, error) {
			return &command.MountCommand{
				Meta: *meta,
			}, nil
		},
		"unmount": func() (cli.Command, error) {
			return &command.UnmountCommand{
				Meta: *meta,
			}, nil
		},
		"waitforattach": func() (cli.Command, error) {
			return &command.WaitForAttachCommand{
				Meta: *meta,
			}, nil
		},
		"getvolumename": func() (cli.Command, error) {
			return &command.GetVolumeNameCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
