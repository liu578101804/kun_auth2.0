package handlers

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/liu578101804/kun_auth2.0/service"
  "github.com/liu578101804/kun_auth2.0/models"
  "github.com/liu578101804/kun_auth2.0/env"
)

func RegUserHandler()  {
  var(
    router *gin.RouterGroup
  )
  router = G_app.Group("/user")

  router.GET("/my/:token", UserPage)

  router.POST("/", UserInfo)
}

func UserPage(c *gin.Context)  {

  c.HTML(http.StatusOK, "home/user", gin.H{})
}

func UserInfo(c *gin.Context)  {
  var(
    token string
    err error
    errCode int
    accessTokenModel *models.OauthAccessToken
    userModel *models.OauthUser
  )
  token = c.DefaultPostForm("token","")

  //验证token
  if accessTokenModel,err = service.CheckToken(token);err != nil {
    errCode = env.ERRCODE_USER_INFO + 1
    goto ERR
  }

  //查询用户信息
  if userModel,err = service.GetUserInfo(accessTokenModel.UserId);err !=nil {
    errCode = env.ERRCODE_USER_INFO + 2
    goto ERR
  }

  c.JSON(http.StatusOK,createSuccess(gin.H{
    "open_id": userModel.OpenId,
    "name": userModel.Name,
    "email": userModel.Email,
  }))

  return
  ERR:
    c.JSON(http.StatusOK, createERR(errCode,err.Error()))
}

