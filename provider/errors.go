package provider

import "errors"

var (
	errInstanceNotExist   = errors.New("instance does not exist")
	errHostnameDuplicated = errors.New("instance hostname duplicated")
	errDiskNotExist       = errors.New("disk does not exist")
)
