// Copyright 2023 jami1024 &lt;996013797@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/jami1024/miniblog.

package model

import (
	"time"

	"github.com/jami1024/miniblog/pkg/auth"
	"gorm.io/gorm"
)

// 利用db2struct 自动生成
// db2struct --gorm --no-json -H 127.0.0.1 -d miniblog -t user --package model --struct UserM -u root -p 'mPew9UlWJjyhrVc46TDp8fXx2' --target=user.go

// UserM 是数据库中 user 记录 struct 格式的映射.
type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"` //
	Username  string    `gorm:"column:username"`       //
	Password  string    `gorm:"column:password"`       //
	Nickname  string    `gorm:"column:nickname"`       //
	Email     string    `gorm:"column:email"`          //
	Phone     string    `gorm:"column:phone"`          //
	CreatedAt time.Time `gorm:"column:createdAt"`      //
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //
}

// TableName 用来指定映射的 MySQL 表名.
func (u *UserM) TableName() string {
	return "user"
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (u *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	// Encrypt the user password.
	u.Password, err = auth.Encrypt(u.Password)
	if err != nil {
		return err
	}

	return nil
}
