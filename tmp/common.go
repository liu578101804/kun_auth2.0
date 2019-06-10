package handler

import (
  "github.com/gin-gonic/gin"
  "github.com/liu578101804/kun_auth2.0/models"
  "time"
  "github.com/liu578101804/kun_auth2.0/service"
  "errors"
)

func createERR(code int,msg string) *gin.H {
  return &gin.H{
    "code": code,
    "msg": msg,
    "data":"",
  }
}

func createSuccess(data interface{}) *gin.H {
  return &gin.H{
    "code": 200,
    "msg": "操作成功",
    "data": data,
  }
}

func checkToken(token string) (*models.OauthAccessToken,error) {
  var(
    err error
    accessTokenModel *models.OauthAccessToken
    nowTime time.Time
  )
  nowTime = time.Now()
  if accessTokenModel,err = service.GetTokenInfo(token);err != nil {
    return nil,err
  }
  //是否过期
  if accessTokenModel.ExpiresAt.Sub(nowTime).Seconds() <= 0 {
    return nil,errors.New("token过期")
  }
  return accessTokenModel,nil
}
