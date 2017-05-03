package conv

import (
	"forum/api/vo"
	"forum/model"
	"forum/services"
	"time"
)

func NewToArticle(form vo.ArticleForm, user model.User) model.Article {
	now := time.Now()
	article := model.Article{
		Title:        services.SensitiveFilter(form.Title),
		Content:      services.SensitiveFilter(form.Content),
		AuthorID:     user.ID,
		AuthorName:   user.Name,
		CreatedOn:    now,
		LastUpdateOn: now,
		State:        model.ArticleChecking,
		ClubID:       form.ClubID,
	}
	return article
}

func UpdateToArticle(form vo.UpdateArticleForm) model.Article {
	article := model.Article{}

	return article
}

func ToArticlesVo(article model.Article) vo.Article {
	return vo.Article{
		ID:            article.ID,
		Title:         article.Title,
		AuthorID:      article.AuthorID,
		AuthorName:    article.AuthorName,
		Abstract:      string([]rune(article.Content)[:50]),
		ContainImage:  false,
		Hits:          article.Hits,
		TotalComments: article.TotalComments,
	}
}
