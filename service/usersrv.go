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
  if userModel.Password != utils.GetPassword(password,userModel.Salt) {
    return nil,errors.New("密码错误")
  }
  return &userModel,nil
}

//注册
func RegisterService(email, password string) (*models.OauthUser, error)  {
  var(
    err error
    regUserModel models.OauthUser
    salt string
    inputPassword string
  )
  inputPassword,salt = utils.CreatePassword(password)
  if _,err = LoginService(email,password);err != nil{
    if err.Error() == "用户未注册" {
      regUserModel = models.OauthUser{
        Email: email,
        Password : inputPassword,
        Salt: salt,
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
func QueryClientByClientKey(clientKey string) (*models.OauthClient,error) {
  var(
    clientModel models.OauthClient
    err error
    has bool
  )
  clientModel = models.OauthClient{
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
func CreateAccessToken(code, clientId string, userId int) (*models.OauthAccessToken,error) {

  var(
    token string
    //refToken string
    nowTime time.Time
    exportAt time.Time
    oatModel models.OauthAccessToken
    err error
    oaCodeModel models.OauthAuthorizationCode
  )

  token = utils.CreateToken()
  //refToken = utils.CreateToken()

  nowTime = time.Now()
  hh, _ := time.ParseDuration(fmt.Sprintf("%vms",config.G_config.TokenExpiresAt))
  exportAt = nowTime.Add(hh)

  //存入toke
  oatModel = models.OauthAccessToken{
    Token: token,
    UserId: userId,
    ClientId: clientId,
    ExpiresAt: exportAt,
  }
  if _,err = database.G_engine.Insert(oatModel);err != nil {
    return nil,err
  }

  //更新code状态
  oaCodeModel = models.OauthAuthorizationCode{
    Code: code,
    IsUse: "T",
  }
  if _,err = database.G_engine.Where("code=?",oaCodeModel.Code).Update(oaCodeModel);err != nil {
    return nil,err
  }

  return &oatModel,nil
}

func GetTokenInfo(token string) (*models.OauthAccessToken, error) {
  var(
    has bool
    accessTokenModel models.OauthAccessToken
    err error
  )
  accessTokenModel = models.OauthAccessToken{
    Token: token,
  }
  if has,err = database.G_engine.Where("token=?",accessTokenModel.Token).Get(&accessTokenModel);err !=nil {
    return nil, err
  }
  if !has {
    err = errors.New("没找到token")
    return nil, err
  }
  return &accessTokenModel, nil
}

func GetUserInfo(userId int) (*models.OauthUser,error) {
  var(
    accessUserModel models.OauthUser
    has bool
    err error
  )
  accessUserModel = models.OauthUser{}
  accessUserModel.BaseModel.Id = userId
  if has,err = database.G_engine.Get(&accessUserModel);err !=nil {
    return nil,err
  }
  if !has {
    return nil,errors.New("没找到token")
  }
  return &accessUserModel,nil
}


func CheckToken(token string) (accessTokenModel *models.OauthAccessToken,err error) {
  var(
    nowTime time.Time
  )
  nowTime = time.Now()
  if accessTokenModel,err = GetTokenInfo(token);err != nil {
    return
  }
  //是否过期
  if accessTokenModel.ExpiresAt.Sub(nowTime).Seconds() <= 0 {
    err = errors.New("token过期")
    return
  }
  return
}
