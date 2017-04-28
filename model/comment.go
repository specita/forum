package model

import "time"

type Comment struct {
	ID         int       `gorm:"AUTO_INCREMENT,PRIMARY_KEY"`
	ArticleID  int       `gorm:"TYPE:INT",NOT NULL,COMMENT '帖子ID'`
	AuthorID   int       `gorm:"TYPE:INT",NOT NULL,COMMENT '作者ID'`
	AuthorName string    `gorm:"NOT NULL,COMMENT '作者姓名'`
	CreatedOn  time.Time `gorm:"NOT NULL,DEFAULT NOW(),COMMENT '创建日期'`
	AtWhoId	   int       `gorm:"TYPE:INT",NOT NULL,COMMENT '回复某人，某人的ID'`
}
