// Copyright 2023 jami1024 &lt;996013797@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/jami1024/miniblog.

package errno

import (
	"errors"
	"fmt"
)

// Errno 定义了moniblog 使用的错误类型
type Errno struct {
	HTTP    int
	Code    string
	Message string
}

// Error 实现了error接口中的`Error`方法
func (err *Errno) Error() string {
	return err.Message
}

// SetMessage 设置Errno 类型错误中的Message字段
func (err *Errno) SetMessage(format string, args ...interface{}) *Errno {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}
	var typed *Errno
	switch {
	case errors.As(err, &typed):
		return typed.HTTP, typed.Code, typed.Message
	default:
		return InternalServerError.HTTP, InternalServerError.Code, InternalServerError.Message
	}
	//return InternalServerError.HTTP, InternalServerError.Code, InternalServerError.Message
}
