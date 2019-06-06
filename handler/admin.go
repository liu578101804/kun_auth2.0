package handler

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func RegAdminHandler()  {
  G_app.GET("/admin", AdminIndexPage)
}

func AdminIndexPage(c *gin.Context)  {

  c.HTML(http.StatusOK, "admin/index", gin.H{
    "title":"blog",
  })

}
