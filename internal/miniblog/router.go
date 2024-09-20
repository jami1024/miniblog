// Copyright 2023 jami1024 &lt;996013797@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/jami1024/miniblog.

package miniblog

import (
	"github.com/gin-gonic/gin"
	"github.com/jami1024/miniblog/internal/miniblog/controller/v1/user"
	"github.com/jami1024/miniblog/internal/miniblog/store"
	"github.com/jami1024/miniblog/internal/pkg/core"
	"github.com/jami1024/miniblog/internal/pkg/errno"
	"github.com/jami1024/miniblog/internal/pkg/log"
	mw "github.com/jami1024/miniblog/internal/pkg/middleware"
	"github.com/jami1024/miniblog/pkg/auth"
)

// installRouters 安装 miniblog 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/health", func(c *gin.Context) {
		log.C(c).Infow("health function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)
	g.POST("/login", uc.Login)
	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			// 创建用户
			userv1.POST("", uc.Create)
			// 修改用户密码
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mw.Authn(), mw.Authz(authz))
			// 获取用户详情
			//userv1.GET(":name", uc.Ge)
		}
	}

	return nil
}
