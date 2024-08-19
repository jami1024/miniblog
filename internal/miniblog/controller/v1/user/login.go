// Copyright 2023 jami1024 &lt;996013797@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/jami1024/miniblog.

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jami1024/miniblog/internal/pkg/core"
	"github.com/jami1024/miniblog/internal/pkg/errno"
	"github.com/jami1024/miniblog/internal/pkg/log"
	v1 "github.com/jami1024/miniblog/pkg/api/miniblog/v1"
)

// Login 登录 miniblog 并返回一个 JWT Token.
func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
