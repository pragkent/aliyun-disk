package provider

import "errors"

var (
	ErrInstanceNotExist   = errors.New("instance does not exist")
	ErrHostnameDuplicated = errors.New("instance hostname duplicated")
	ErrDiskNotExist       = errors.New("disk does not exist")
	ErrDiskTagDuplicated  = errors.New("disk tags duplicated")
)
