package model

import "time"

//评论
type Comment struct {
	ID         int       `gorm:"AUTO_INCREMENT,PRIMARY_KEY"`
	ArticleID  int       `gorm:"NOT NULL;COMMENT:'帖子ID'"`
	AuthorID   string    `gorm:"NOT NULL;COMMENT:'作者ID'"`
	AuthorName string    `gorm:"NOT NULL;COMMENT:'作者姓名'"`
	CreatedOn  time.Time `gorm:"NOT NULL;DEFAULT:NOW();COMMENT:'创建日期'"`
	AtWhoID    int       `gorm:"NOT NULL;COMMENT:'被回复人ID'"`
	AtWhoName  string    `gorm:"NOT NULL;COMMENT:'被回复人名'"`
	Content    string    `gorm:"NOT NULL;COMMENT:'回复内容'"`
	UpVote     int       `gorm:"DEFAULT:0;COMMENT:'顶'"`
	DownVote   int       `gorm:"DEFAULT:0;COMMENT:'踩'"`
}
