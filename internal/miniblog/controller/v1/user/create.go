// Copyright 2023 jami1024 &lt;996013797@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/jami1024/miniblog.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/jami1024/miniblog/internal/pkg/core"
	"github.com/jami1024/miniblog/internal/pkg/errno"
	"github.com/jami1024/miniblog/internal/pkg/log"
	v1 "github.com/jami1024/miniblog/pkg/api/miniblog/v1"
)

func (ctrl *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")
	var r v1.CreateUserRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	// govalidator 参数校验
	// https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fasaskevich%2Fgovalidator
	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := ctrl.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}
	core.WriteResponse(c, nil, nil)
}
