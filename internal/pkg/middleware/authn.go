// Copyright 2023 jami1024 &lt;996013797@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/jami1024/miniblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jami1024/miniblog/internal/pkg/core"
	"github.com/jami1024/miniblog/internal/pkg/errno"
	"github.com/jami1024/miniblog/internal/pkg/known"
	"github.com/jami1024/miniblog/pkg/token"
)

// Authn 是认证中间件，用来从 gin.Context 中提取 token 并验证 token 是否合法，
// Authz 是鉴权
// 如果合法则将 token 中的 sub 作为<用户名>存放在 gin.Context 的 XUsernameKey 键中.

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT token
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()

			return
		}
		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}
