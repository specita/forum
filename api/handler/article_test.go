package handler

import (
	"forum/api/vo"
	"forum/conv"
	"forum/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	ArticleForm vo.ArticleForm
	CurrentUser model.User
)

type ArticleSuite struct {
	suite.Suite
	Err error
}

func (as *ArticleSuite) SetupSuite() {
	ArticleForm.Title = "第一个帖子"
	ArticleForm.Content = "求赞求经验啊。。。。。。。"
	ArticleForm.Images = []string{"http://www.baidu.com", "http://www.sina.com"}
	ArticleForm.BlackUserIDs = []string{"xxx-user1", "xxx-user2"}

	CurrentUser.ID = "xxx-user3"
	CurrentUser.Name = "pchen"

	model.InitDB("root:coolmoon1@(127.0.0.1:3306)/forum?charset=utf8&parseTime=True&loc=Local", true)
}

func (as *ArticleSuite) TestNewArticleInDB() {
	article := conv.NewToArticle(ArticleForm, CurrentUser)
	as.Err = newArticleInDB(article, ArticleForm.BlackUserIDs, ArticleForm.Images)
}

func TestRunSuite(t *testing.T) {
	s := new(ArticleSuite)
	suite.Run(t, s)

	require.NoError(t, s.Err, "insert article failed!")
}
