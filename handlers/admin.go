package handlers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func RegAdminHandler()  {
  var(
    router *gin.RouterGroup
  )
  router = G_app.Group("/admin")

  router.GET("/", AdminIndexPage)
}

func AdminIndexPage(c *gin.Context)  {

  c.HTML(http.StatusOK, "admin/index", gin.H{
    "title":"blog",
  })

}
