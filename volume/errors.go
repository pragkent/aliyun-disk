package volume

import (
	"errors"
)

var (
	errAliyunAPIUnavailable = errors.New("aliyun API is unavailable")
	errInvalidVolumeOptions = errors.New("invalid volume options")
	errCommandNotSupported  = errors.New("command not supported")
)

func NewDriverError(err error) *DriverStatus {
	return &DriverStatus{
		Status:  StatusFailure,
		Message: err.Error(),
	}
}

func NewDriverNotSupported(err error) *DriverStatus {
	return &DriverStatus{
		Status:  StatusNotSupported,
		Message: err.Error(),
	}
}
