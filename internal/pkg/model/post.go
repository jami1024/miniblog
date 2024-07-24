package model

import "time"

// PostM 是数据库中 post 记录 struct 格式的映射.
// 通过db2struct --gorm --no-json -H 127.0.0.1 -d miniblog -t post --package model --struct PostM -u root -p 'mPew9UlWJjyhrVc46TDp8fXx2' --target=post.go
type PostM struct {
	ID        int64     `gorm:"column:id;primary_key"` //
	Username  string    `gorm:"column:username"`       //
	PostID    string    `gorm:"column:postID"`         //
	Title     string    `gorm:"column:title"`          //
	Content   string    `gorm:"column:content"`        //
	CreatedAt time.Time `gorm:"column:createdAt"`      //
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //
}

// TableName 用来指定映射的 MySQL 表名.
func (p *PostM) TableName() string {
	return "post"
}
