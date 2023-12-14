package xcode

import (
	"context"
	"github.com/pkg/errors"
)

func CodeFromError(err error) XCode {
	err = errors.Cause(err)
	if code, ok := err.(XCode); ok {
		return code
	}

	switch err {
	case context.Canceled:
		return Canceled
	case context.DeadlineExceeded:
		return Deadline
	default:
		return ServerErr
	}
}
