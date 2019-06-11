package handlers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func RegHomeHandler() {
  G_app.GET("/", HomePage)
}

func HomePage(c *gin.Context) {
  c.HTML(http.StatusOK,"home/index", gin.H{
    "title":"首页",
  })
}

