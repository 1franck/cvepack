package core

import "fmt"

type ErrorMsg string

const (
	EmptyError ErrorMsg = ""
)

func (e ErrorMsg) ToString() string {
	return string(e)
}

func (e ErrorMsg) Sprintf(args ...interface{}) ErrorMsg {
	return ErrorMsg(fmt.Sprintf(e.ToString(), args...))
}
