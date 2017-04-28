package handler

import (
	//"github.com/go-martini/martini"
	"forum/api/middleware/requestid"
	"forum/api/vo"
	"forum/conv"
	"forum/model"
	"github.com/Sirupsen/logrus"
	"github.com/martini-contrib/render"
	"net/http"
)

func NewArticle(form vo.ArticleForm, reqID requestid.RequestId, logger logrus.Entry, r render.Render) {
	article := conv.ToArticle(form)

	if err := model.DB.Create(&article).Error; err != nil {
		logger.Error(err)
		r.JSON(http.StatusInternalServerError, "")
		return
	}

	r.JSON(http.StatusOK, "ok")
}

func DeleteArticle(form vo.DelArticleForm, reqID requestid.RequestId, logger logrus.Entry, r render.Render) {
	//hasDelPermisson
}
