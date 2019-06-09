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

//注册
func RegisterService(email, password string) (*models.OauthUser, error)  {
  var(
    err error
    regUserModel models.OauthUser
  )
  if _,err = LoginService(email,password);err != nil{
    if err.Error() == "用户未注册" {
      regUserModel = models.OauthUser{
        Email: email,
        Password : password,
        OpenId: utils.CreateUserOpenId(),
      }
      if _,err = database.G_engine.InsertOne(&regUserModel);err != nil {
        goto ERR
      }
      return &regUserModel,nil
    }else{
      err = errors.New("用户已经存在")
      goto ERR
    }
  }
  err = errors.New("用户已经存在")
ERR:
  return nil,err
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
    nowTime time.Time
    exportAt time.Time
  )

  nowTime = time.Now()
  hh, _ := time.ParseDuration(fmt.Sprintf("%vms",config.G_config.CodeExpiresAt))
  exportAt = nowTime.Add(hh)

  oacModel = models.OauthAuthorizationCode{
    Code: code,
    UserId: UserId,
    ClientId: ClientId,
    ExpiresAt: exportAt,
    IsUse: "F",
  }
  if _,err = database.G_engine.Insert(oacModel);err != nil {
    return
  }
  return
}

// 通过code查到相关信息
func GetAuthorizationCodeByCode(code string) (*models.OauthAuthorizationCode,error) {
  var(
    oacModel models.OauthAuthorizationCode
    has bool
    err error
  )
  oacModel = models.OauthAuthorizationCode{
    Code: code,
  }
  if has,err = database.G_engine.Get(&oacModel);err != nil {
    return nil, err
  }
  if !has {
    return nil, errors.New("没找到code")
  }
  return &oacModel, nil
}

// 生成Token
func CreateAccessToken(clientId string, userId int) (*models.OauthAccessToken,error) {

  var(
    token string
    //refToken string
    nowTime time.Time
    exportAt time.Time
    oatModel models.OauthAccessToken
    err error
  )

  token = utils.CreateToken()
  //refToken = utils.CreateToken()

  nowTime = time.Now()
  hh, _ := time.ParseDuration(fmt.Sprintf("%vms",config.G_config.TokenExpiresAt))
  exportAt = nowTime.Add(hh)

  oatModel = models.OauthAccessToken{
    Token: token,
    UserId: userId,
    ClientId: clientId,
    ExpiresAt: exportAt,
  }
  if _,err = database.G_engine.Insert(oatModel);err != nil {
    return nil,err
  }

  return &oatModel,nil
}


//获取用户授权的结构
