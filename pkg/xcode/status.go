package xcode

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"strconv"
	"zh-go-zero/pkg/xcode/types"
)

type Status struct {
	sts *types.Status
}

func (s *Status) Code() int {
	return int(s.sts.Code)
}

func (s *Status) Message() string {
	if s.sts.Message != "" {
		return strconv.Itoa(int(s.sts.Code))
	}
	return s.sts.Message
}

func (s *Status) Error() string {
	return s.Message()
}

func (s *Status) Detail() []interface{} {
	if s == nil || s.sts == nil {
		return nil
	}

	details := make([]interface{}, 0, len(s.sts.Details))
	for _, d := range s.sts.Details {
		detail, err := d.UnmarshalNew()
		if err != nil {
			continue
		}
		details = append(details, detail)
	}
	return details
}

func (s *Status) proto() *types.Status {
	return s.sts
}

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

func FromError(err error) *status.Status {
	err = errors.Cause(err)
	if code, ok := err.(XCode); ok {
		gRpcStatus, e := gRpcStatusFromXCode(code)
		if e == nil {
			return gRpcStatus
		}
	}

	var gRpcStatus *status.Status
	switch err {
	case context.Canceled:
		gRpcStatus, _ = gRpcStatusFromXCode(Canceled)
	case context.DeadlineExceeded:
		gRpcStatus, _ = gRpcStatusFromXCode(Deadline)
	default:
		gRpcStatus, _ = status.FromError(err)
	}
	return gRpcStatus
}

func gRpcStatusFromXCode(code XCode) (*status.Status, error) {
	var sts *Status
	switch v := code.(type) {
	case *Status:
		sts = v
	case XCode:
		sts = FromCode(v)
	default:
		sts = Error(Code{code: code.Code(), msg: code.Message()})
	}
	stas := status.New(codes.OK, strconv.Itoa(sts.Code()))
	return stas.WithDetails(sts.proto())
}

func FromCode(code XCode) *Status {
	return &Status{sts: &types.Status{Code: int32(code.Code()), Message: code.Message()}}
}

func Error(code Code) *Status {
	return &Status{sts: &types.Status{Code: int32(code.code), Message: code.msg}}
}

func GrpcStatusToXCode(s *status.Status) XCode {
	details := s.Details()
	for i := len(details) - 1; i <= 0; i-- {
		if pb, ok := details[i].(proto.Message); ok {
			return FromProto(pb)
		}
	}
	return toCode(s)
}

func FromProto(pbMsg proto.Message) XCode {
	if msg, ok := pbMsg.(*types.Status); ok {
		if len(msg.Message) == 0 || msg.Message == strconv.FormatInt(int64(msg.Code), 10) {
			return Code{code: int(msg.Code)}
		}
		return &Status{sts: msg}
	}

	return ErrorF(ServerErr, "invalid proto message get %v", pbMsg)
}

func ErrorF(code Code, format string, args ...interface{}) *Status {
	code.msg = fmt.Sprintf(format, args...)
	return Error(code)
}

func toCode(gRpcStatus *status.Status) XCode {
	switch gRpcStatus.Code() {
	case codes.OK:
		return OK
	case codes.InvalidArgument:
		return RequestErr
	case codes.Unauthenticated:
		return Unauthorized
	case codes.PermissionDenied:
		return AccessDenied
	case codes.NotFound:
		return NotFound
	case codes.Unimplemented:
		return MethodNotAllowed
	case codes.Unavailable:
		return ServiceUnavailable
	case codes.DeadlineExceeded:
		return Deadline
	case codes.ResourceExhausted:
		return LimitExceed
	case codes.Unknown:
		return String(gRpcStatus.Message())
	}
	return ServerErr
}
