package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (in IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"title": "Main website",
	})
}
