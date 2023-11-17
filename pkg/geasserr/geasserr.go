/*
----------------------------------------
@Create 2023/11/16
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe err 15:36
----------------------------------------
@Version 1.0 2023/11/16
@Memo create this file
*/

package geasserr

import (
	"errors"
	"fmt"
)

type Code int

func (c Code) New(args ...any) Error {
	geasserr := NewGeasserr(c)
	geasserr.Message = fmt.Sprintf("%v", args)
	return geasserr
}

type Error struct {
	Code     Code   `json:"code"`
	Message  string `json:"message"`
	Producer string `json:"producer"`
}

func (g Error) Error() string {
	return fmt.Sprintf("[%s-%d]%s", g.Producer, g.Code, g.Message)
}

func NewGeasserr(code Code) Error {
	return Error{
		Code:     code,
		Message:  "",
		Producer: "",
	}
}

func CheckErrCode(err error) Code {
	var geasserr Error
	if errors.As(err, &geasserr) {
		return geasserr.Code
	}
	return -1
}
