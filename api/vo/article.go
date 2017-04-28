package vo

import (
	"time"
)

type ArticleForm struct {
	Title        string    `json:"title"`
	AuthorID     int       `json:"authodID"`
	AuthorName   string    `json:"authorName"`
	CreatedOn    time.Time `json:"createdOn"`
	LastUpdateOn time.Time `json:"lastUpdateOn"`
	BlackUserIDs []int     `json:"blackUsersIDs"`
	Content      string    `json:"content"`
	//图片压缩包路径,TODO 手机端可以设置压缩，但网页端？？
	Images []string `json:"images"`
}

type DelArticleForm struct {
	ArticleID  string `json:"articleID" binding:"required"`
	OperatorID string `json:"operatorID" binding:"required"`
}
