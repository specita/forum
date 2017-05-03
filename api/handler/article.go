package handler

import (
	"forum/api/middleware/requestid"
	"forum/api/vo"
	"forum/common"
	"forum/conv"
	"forum/model"
	"forum/services"
	"github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	"strconv"
	"time"
)

//新增帖子
func NewArticle(form vo.ArticleForm, reqID requestid.RequestId, log *log.Logger, r render.Render, user *model.User) {
	article := conv.NewToArticle(form, *user)

	if err := newArticleInDB(article, form.BlackUserIDs, form.Images); err != nil {
		log.Println(err)
		r.JSON(http.StatusInternalServerError, "")
		return
	}

	r.JSON(http.StatusCreated, "ok")
	return
}

//新贴入库
func newArticleInDB(article model.Article, blackUserIDs []string, images []string) error {
	tx := model.DB.Begin()

	if len(images) > 0 {
		article.ContainImages = true

	}

	//保存帖子
	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		return err
	}

	//保存帖子黑名单
	if len(blackUserIDs) > 0 {
		blackList := [][]interface{}{}
		for _, userID := range blackUserIDs {
			blackList = append(blackList, []interface{}{article.ID, article.AuthorID, userID})
		}
		//
		sql, vals := common.BatchInsert("article_blacklist", []string{"article_id", "author_id", "black_user_id"}, blackList)
		if _, err := tx.CommonDB().Exec(sql, vals...); err != nil {
			tx.Rollback()
			return err
		}
	}

	//保存图像链接
	if len(images) > 0 {
		imageList := [][]interface{}{}
		for _, image := range images {
			imageList = append(imageList, []interface{}{article.ID, image})
		}

		sql, vals := common.BatchInsert("article_images", []string{"article_id", "url"}, imageList)
		if _, err := tx.CommonDB().Exec(sql, vals...); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

//删贴
func DeleteArticle(form vo.DelArticleForm, reqID requestid.RequestId, logger logrus.Entry, r render.Render, user model.User) {
	//todo hasDelPermisson??
	article := model.Article{ID: form.ArticleID}

	if err := model.DB.First(&article).Error; err != nil {
		logger.Error(err)
		r.JSON(http.StatusNotFound, "no this article")
		return
	}

	deleted := model.DeletedArticle{
		ArticleID:        article.ID,
		AuthorID:         article.AuthorID,
		DeletedOn:        time.Now(),
		ArticleCreatedOn: article.CreatedOn,
		Title:            article.Title,
		OperatorID:       user.ID,
	}

	tx := model.DB.Begin()

	if err := model.DB.Create(&deleted).Error; err != nil {
		tx.Rollback()
		return
	}

	if err := model.DB.Delete(&article).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	r.JSON(http.StatusOK, "true")
}

//更新帖子
func UpdateArticle(form vo.UpdateArticleForm, logger logrus.Entry, reqID requestid.RequestId, tx *gorm.DB, r render.Render) {
	article := conv.UpdateToArticle(form)
	article.Content = services.SensitiveFilter(form.Content)

	tx.Model(&article).Updates(model.Article{
		Title:        article.Title,
		Content:      article.Content,
		LastUpdateOn: time.Now(),
	})

	image := model.ArticleImages{ArticleID: article.ID}
	for _, del := range form.DeleteImages {
		tx.Where("url = ?", del).Delete(image)
	}

	//todo batch insert upload images
	r.JSON(http.StatusOK, "update ok")
}

//帖子列表
func AllArticles(form vo.ArticlesForm, logger *log.Logger, r render.Render) {
	form.Pagination.CheckOrResetBounds()
	articles := []model.Article{}
	if err := model.DB.Where("club_id = ? AND state = ?", form.ClubID, model.ArticleChecked).
		Order("sort_id asc, created_on desc").Limit(form.Limit).Offset(form.Offset).Find(&articles).Error; err != nil {
		logger.Println(err)
		r.JSON(http.StatusInternalServerError, "")
		return
	}

	result := make([]vo.Article, 0, len(articles))
	for _, a := range articles {
		result = append(result, conv.ToArticlesVo(a))
	}

	r.JSON(http.StatusOK, result)
}

//审核帖子
func ArticlePass(params martini.Params, log *log.Logger, r render.Render) {
	articleID, _ := strconv.Atoi(params["ID"])

	article := model.Article{ID: articleID}
	if err := model.DB.Model(&article).UpdateColumn("state", model.ArticleChecked).Error; err != nil {
		log.Println(err)
		r.JSON(http.StatusInternalServerError, "db err")
		return
	}

	r.JSON(http.StatusOK, "ok")
}

//进入帖子，获取详情
func ArticleDetail(params martini.Params, log *log.Logger, r render.Render) {
	//articleID, _ := strconv.Atoi(params["ID"])

}
