package service

import (
  "github.com/liu578101804/kun_auth2.0/database"
  "github.com/liu578101804/kun_auth2.0/models"
  "errors"
  "fmt"
  "github.com/liu578101804/kun_auth2.0/config"
  "time"
  "github.com/liu578101804/kun_auth2.0/utils"
)

// 登录
func LoginService(email, password string) (*models.OauthUser, error) {
  var(
    userModel models.OauthUser
    err error
    has bool
  )
  userModel = models.OauthUser{
    Email: email,
  }
  //查找用户
  if has,err= database.G_engine.Get(&userModel); err != nil{
    return nil,err
  }
  //是否存在
  if !has {
    return nil,errors.New("用户未注册")
  }
  //检查密码是否正确
  if userModel.Password != password {
    return nil,errors.New("密码错误")
  }
  return &userModel,nil
}

// 查询client信息
func QueryClientByClientKey(clientKey string) (*models.OauthClients,error) {
  var(
    clientModel models.OauthClients
    err error
    has bool
  )
  clientModel = models.OauthClients{
    ClientKey: clientKey,
  }
  if has,err = database.G_engine.Get(&clientModel);err != nil {
    return nil,err
  }
  if !has {
    return nil,errors.New("没找到数据")
  }
  return &clientModel,nil
}

// 生成登录后的code
func CreateCode(UserId int, ClientId, code string) (err error) {
  var(
    oacModel models.OauthAuthorizationCode
  )
  oacModel = models.OauthAuthorizationCode{
    Code: code,
    UserId: UserId,
    ClientId: ClientId,
  }
  if _,err = database.G_engine.Insert(oacModel);err != nil {
    return
  }
  return
}

// 生成Token
func CreateAccessToken(clientId string, userId int) (err error) {

  var(
    token string
    refToken string
    nowTime time.Time
    exportAt time.Time
  )

  token = utils.CreateToken()
  refToken = utils.CreateToken()

  nowTime = time.Now()
  hh, _ := time.ParseDuration(fmt.Sprintf("%vms",config.G_config.TokenExpiresAt))
  exportAt = nowTime.Add(hh)

  var(
    oatModel models.OauthAccessToken
  )
  oatModel = models.OauthAccessToken{
    Token: token,
    UserId: userId,
    ClientId: clientId,
    ExpiresAt: exportAt,
  }
  if _,err = database.G_engine.Insert(oatModel);err != nil {
    return
  }
  return



  exportAt = exportAt
  refToken = refToken

  return
}
