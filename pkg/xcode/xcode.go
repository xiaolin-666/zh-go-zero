package xcode

import "strconv"

type Code struct {
	code int
	msg  string
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
