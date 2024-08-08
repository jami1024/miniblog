package biz

import "github.com/jami1024/miniblog/internal/miniblog/store"
import "github.com/jami1024/miniblog/internal/miniblog/biz/user"

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

func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}
