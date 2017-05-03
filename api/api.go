package api

import (
	"forum/api/middleware/authorized"
	"forum/api/middleware/logrus"
	"forum/api/middleware/requestid"
	"forum/config"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Run() {
	m := forumMartini()
	//m := martini.Classic()
	m.Head("/healthz", func() string { return "ok" })
	m.Get("/healthz", func() string { return "ok" })

	m.Group("/api", func(r martini.Router) {
		r.Group("/auth", AuthRouter)
		r.Group("/articles", ArticleRouter, authorized.Authorized())
		r.Group("/comments", CommentRouter)
		//r.Group("/upvote", UpVoteRouter)
	})

	m.RunOnAddr(config.Config.Address)
}

//init martini
func forumMartini() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()

	//martini.Env = martini.Prod

	m.Use(logrus.Logger())
	m.Use(martini.Recovery())
	m.Use(render.Renderer())
	m.Use(requestid.Middleware(&requestid.Options{true}))
	m.Use(martini.Static("html"))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	return &martini.ClassicMartini{m, r}
}
