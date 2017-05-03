package authorized

import (
	"forum/model"
	"github.com/go-martini/martini"
)

func Authorized() martini.Handler {
	return func(c martini.Context) {
		user := model.User{"xxx-user1", "pchen", "http://ivs.com"}
		c.Map(&user)
	}
}
