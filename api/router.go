package api

import (
	"forum/api/handler"
	"forum/api/vo"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

//用户登录注册相关
func AuthRouter(r martini.Router) {
	r.Post("/register", func() string {
		return "ok"
	})
	r.Post("/login", func() string {
		return "ok"
	})
	r.Post("/request_sms_code", func() string {
		return "ok"
	})
	r.Post("/logout", func() string {
		return "ok"
	})
}

//管理员相关操作
func AdminRouter(r martini.Router) {
	r.Post("/")
}

//帖子相关
func ArticleRouter(r martini.Router) {
	//发贴
	r.Post("/new", binding.Bind(vo.ArticleForm{}), handler.NewArticle)

	//删贴
	r.Post("/delete", binding.Bind(vo.DelArticleForm{}), handler.DeleteArticle)

	//更新
	r.Post("/update")

	//搜索
	r.Get("/search")

	//展现所有
	r.Get("/list")
}

//评论相关
func CommentRouter(r martini.Router) {
	//给帖子发表评论
	r.Post("/new:artcile_id", binding.Bind(vo.ArticleForm{}), handler.NewArticle)

	//回复某人
	r.Post("/reply")

	//删除
	r.Post("/delete")
}

//点赞相关
func UpVoteRouter(r martini.Router) {
	//点赞 or 取消赞
	r.Post("/toggle/:article_id")
}
