package model

import "time"

//帖子状态，审核中和已审核
const (
	ArticleChecking = iota
	ArticleChecked
)

//主题
//todo gorm不能解析COMMENT字段??
type Article struct {
	ID            int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Title         string    `gorm:"NOT NULL;COMMENT:'标题'"`
	AuthorID      string    `gorm:"NOT NULL;COMMENT:'作者id'"`
	ClubID        int       `gorm:"NOT NULL;index:idx_name_club;COMMENT:'企业ID'"`
	AuthorName    string    `gorm:"NOT NULL;COMMENT:'作者名称'"`
	CreatedOn     time.Time `gorm:"NOT NULL;COMMENT:'创建时间'"`
	LastUpdateOn  time.Time `gorm:"NOT NULL;DEFAULT:NOW();COMMENT:'最后更新时间'"`
	Content       string    `gorm:"NOT NULL"`
	Hits          int       `gorm:"NOT NULL;DEFAULT:0;COMMENT:'点击次数'"`
	TotalComments int       `gorm:"NOT NULL;DEFAULT:0;COMMENT:'评论总数'"`
	ContainImages bool      `gorm:"NOT NULL;DEFAULT:0,COMMENT:'是否包含图片(冗余)'"`
	State         int       `gorm:"NOT NULL;DEFAULT:0;COMMENT:'状态，审核中=0，已审核=1'`
	SortID        int       `gorm:"NOT NULL;DEFAULT:99;COMMENT:'排序数字'"`
}

//软删除的主题
type DeletedArticle struct {
	ID               int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	ArticleID        int       `gorm:"NOT NULL;COMMENT:'帖子ID'"`
	Title            string    `gorm:"NOT NULL;COMMENT:'标题'"`
	AuthorID         string    `gorm:"NOT NULL;COMMENT:'作者id'"`
	ArticleCreatedOn time.Time `gorm:"NOT NULL;COMMENT:'帖子创建时间'"`
	DeletedOn        time.Time `gorm:"NOT NULL;DEFAULT:NOW();COMMENT:'删除时间'"`
	OperatorID       string    `gorm:NOT NULL;COMMENT:'操作者ID'`
}

//帖子哪些人不可见
type ArticleBlacklist struct {
	ID          int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	ArticleID   int    `gorm:"NOT NULL;COMMENT:'贴子ID'"`
	AuthorID    string `gorm:"NOT NULL;COMMENT:'作者ID'"`
	BlackUserID string `gorm:"NOT NULL;COMMENT:'不可见用户ID'"`
}

//帖子附图
type ArticleImages struct {
	ID        int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	ArticleID int    `gorm:"NOT NULL;COMMENT:'贴子ID'"`
	Url       string `gorm:"NOT NULL;COMMENT:'图片链接'"`
}

//贴子关注者
type ArticleNotes struct {
	ID         int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	ArticleID  int    `gorm:"NOT NULL;COMMENT:'贴子ID'"`
	ObserverID string `gorm:"NOT NULL;COMMENT:'关注者ID'"`
}
