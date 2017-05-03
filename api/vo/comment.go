package vo

import "time"

type Comment struct {
	ID         int       `json:"id"`
	ArticleID  int       `json:"article_id"`
	AuthorID   string    `json:"author_id"`
	AuthorName string    `json:"author_name"`
	CreatedOn  time.Time `json:"created_on"`
	AtWhoID    int       `json:"at_who_id"`
	AtWhoName  string    `json:"at_who_name"`
	Content    string    `json:"content"`
	UpVote     int       `json:"up_vote"`
	DownVote   int       `json:"down_vote"`
}
