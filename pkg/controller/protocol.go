package controller

import "fmt"

type JsonResult struct {
	Code StatusCode `json:"code"`
	Msg  string     `json:"msg"`
}

type StatusCode int

const (
	StatusCodeOK StatusCode = iota
	StatusCodeFail
	StatusCodeNotFound
	StatusCodeErrorDataBase
	StatusCodeErrorPassword
)

func (c StatusCode) String() string {
	switch c {
	case StatusCodeOK:
		return "StatusCodeOK"
	case StatusCodeFail:
		return "StatusCodeFail"
	default:
		return fmt.Sprintf("error code: %d", c)
	}
}
