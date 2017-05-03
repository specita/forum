package vo

//用于列表展现
type Article struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	AuthorID   string `json:"author_id" binding:"required"`
	AuthorName string `json:"author_name"`
	//摘要
	Abstract      string `json:"abstract"`
	ContainImage  bool   `json:"contain_image"`
	Hits          int    `json:"hits"`
	TotalComments int    `json:"total_comments"`
}

//用于进入帖子获取详情
type ArticleDetail struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	AuthorID      string    `json:"author_id" binding:"required"`
	AuthorName    string    `json:"author_name"`
	Content       string    `json:"content"`
	Images        []string  `json:"images"`
	Hits          int       `json:"hits"`
	TotalComments int       `json:"total_comments"`
	Comments      []Comment `json:"comments"`
}
