package model

//点赞
type UpVote struct {
	ID        int
	ArticleID int `gorm:"index:idx_name_article"`
	UpVoterID int
}
