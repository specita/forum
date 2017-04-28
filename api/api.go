package api

import (
	"forum/api/middleware/requestid"
	"forum/api/middleware/logrus"
	"forum/config"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Run() {
	m := martini.Classic()

	//martini.Env = martini.Prod
	m.Use(render.Renderer())
	m.Use(requestid.Middleware(&requestid.Options{true}))
	m.Handlers(logrus.Logger())



	m.Head("/healthz", func() string { return "ok" })
	m.Get("/healthz", func() string { return "ok" })

	m.Group("/api", func(r martini.Router) {
		r.Group("/auth", AuthRouter)
		r.Group("/articles", ArticleRouter)
		//r.Group("/comments", CommentRouter)
		//r.Group("/upvote", UpVoteRouter)
	})

	m.Get("/hello", func() string {
		return "hello world"
	})
	m.RunOnAddr(config.Config.Address)
}
