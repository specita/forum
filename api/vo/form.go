package vo

type ArticleForm struct {
	Title  string `json:"title" binding:"required"`
	ClubID int    `json:"ClubID" binding:"required"`
	//AuthorID     string   `json:"authodID" binding:"required"`
	//AuthorName   string   `json:"authorName"`
	BlackUserIDs []string `json:"blackUserIDs"`
	Content      string   `json:"content"`
	//图片上传路径
	Images []string `json:"images"`
}

type DelArticleForm struct {
	ArticleID  int    `json:"articleID" binding:"required"`
	OperatorID string `json:"operatorID" binding:"required"`
}

type UpdateArticleForm struct {
	ID           int      `json:"artcileID" binding:"required"`
	Title        string   `json:"title" binding:"required"`
	BlackUserIDs []string `json:"blackUsersIDs" binding:"required"`
	Content      string   `json:"content" binding:"required"`
	//新增图片上传路径
	UploadImages []string `json:"upImages"`
	//删除的图片路径
	DeleteImages []string `json:"delImages"`
}

type ArticlesForm struct {
	ClubID int `json:"clubID" form:"clubID" bingding:"required"`
	Pagination
}
