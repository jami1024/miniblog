package user

import (
	"context"
	"regexp"

	"github.com/jami1024/miniblog/internal/miniblog/store"
	"github.com/jami1024/miniblog/internal/pkg/errno"
	"github.com/jami1024/miniblog/internal/pkg/model"
	"github.com/jinzhu/copier"
)
import v1 "github.com/jami1024/miniblog/pkg/api/miniblog/v1"

// UserBiz 定义了 user 模块在 biz 层所实现的方法.
type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
}

// UserBiz 接口的实现.
type userBiz struct {
	ds store.IStore
}

// 确保 userBiz 实现了 UserBiz 接口.
var _UserBiz = (*userBiz)(nil)

// New 创建一个实现了 UserBiz 接口的实例.
func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}

// Create 是 UserBiz 接口中 `Create` 方法的实现.
func (b *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	// copier 是一个不错的golang库 https://darjun.github.io/2020/03/13/godailylib/copier/
	_ = copier.Copy(&userM, r)

	err := b.ds.Users().Create(ctx, &userM)
	if err != nil {
		match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error())
		if match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}
	return nil
}
