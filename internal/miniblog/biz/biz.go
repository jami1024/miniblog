package biz

import (
	"github.com/jami1024/miniblog/internal/miniblog/biz/user"
	"github.com/jami1024/miniblog/internal/miniblog/store"
)

// IBiz 定义Biz层需要实现的方法
type IBiz interface {
	Users() user.UserBiz
}

// biz 是IBiz的一个实现
type biz struct {
	ds store.IStore
}

// 确保biz实现了IBiz接口
var _IBiz = (*biz)(nil)

// NewBiz 创建
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}
