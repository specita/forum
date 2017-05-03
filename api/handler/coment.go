package handler

import (
	"github.com/gin-gonic/gin/render"
	"github.com/go-martini/martini"
	"log"
	"strconv"
)

func NewComment(params martini.Params, r render.Render, log *log.Logger) {
	articleID, _ := strconv.Atoi(params["ID"])


}
