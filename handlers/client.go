package handlers

import (
  "net/http"
  "errors"
  "github.com/gin-gonic/gin"
  "github.com/liu578101804/kun_auth2.0/service"
  "github.com/liu578101804/kun_auth2.0/models"
  "github.com/liu578101804/kun_auth2.0/env"
)

func RegClient()  {
  var(
    router *gin.RouterGroup
  )
  router = G_app.Group("/client")

  router.POST("/add", clientsAdd)
}

func clientsAdd(c *gin.Context)  {
  var(
    appName string
    redirectUrl string
    token string
    authClientM *models.OauthClient
    af int64
    err error
    errCode int
    accessTokenM *models.OauthAccessToken
  )

  token = c.DefaultPostForm("token","")
  appName = c.DefaultPostForm("app_name","")
  redirectUrl = c.DefaultPostForm("redirect_url","")

  //验证token
  if accessTokenM,err = service.GetTokenInfo(token);err != nil {
    errCode = env.ERRCODE_CLIENT_ADD + 1
    goto ERR
  }

  authClientM = service.CreateClient()
  authClientM.AppName = appName
  authClientM.RedirectUrl = redirectUrl
  authClientM.UserId = accessTokenM.UserId

  if af,err = service.ClientAddService(authClientM);err != nil {
    errCode = env.ERRCODE_CLIENT_ADD + 2
    goto ERR
  }

  if af <= 0 {
    errCode = env.ERRCODE_CLIENT_ADD + 3
    err = errors.New("系统异常")
    goto ERR
  }

  c.JSON(http.StatusOK, createSuccess(gin.H{
    "id": authClientM.BaseModel.Id,
  }))

  return
  ERR:
    c.JSON(http.StatusOK,createERR(errCode, err.Error()))
}
