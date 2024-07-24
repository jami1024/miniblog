// Copyright 2023 jami1024 &lt;996013797@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/jami1024/miniblog.

package errno

var (
	// OK 代表请求成功
	OK = &Errno{HTTP: 200, Code: "", Message: ""}

	// InternalServerError 表示所有的未知的服务器端错误
	InternalServerError = &Errno{
		HTTP:    500,
		Code:    "InternalError",
		Message: "Internal server error",
	}

	// ErrPageNotFound 表示路由不匹配
	ErrPageNotFound = &Errno{
		HTTP:    404,
		Code:    "ResourceNotFound.PageNotFound",
		Message: "Page not found.",
	}
)
