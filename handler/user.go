package handler

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func RegUserHandler()  {

  G_app.GET("/user/:token", UserPage)
  
}

func UserPage(c *gin.Context)  {
  var(
    token string
  )
  token = c.Param("token")

  //验证token
  token = token

  //查询用户信息

  //返回用户信息页面
  c.HTML(http.StatusOK, "home/user", gin.H{})
}
