package xcode

import "strconv"

type XCode interface {
	Code() int
	Message() string
	Detail() []interface{}
	Error() string
}

type Code struct {
	code int
	msg  string
}

func (c Code) Code() int {
	return c.code
}

func (c Code) Message() string {
	return c.Error()
}

func (c Code) Detail() []interface{} {
	return nil
}

func (c Code) Error() string {
	if len(c.msg) > 0 {
		return c.msg
	}
	return strconv.Itoa(c.code)
}

func New(code int, msg string) Code {
	return Code{code: code, msg: msg}
}

func add(code int, msg string) Code {
	return Code{code: code, msg: msg}
}
