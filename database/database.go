package database

import (
  "github.com/go-xorm/xorm"
  "github.com/liu578101804/kun_auth2.0/config"
  "github.com/liu578101804/kun_auth2.0/models"
  "github.com/liu578101804/kun_auth2.0/utils"
  _ "github.com/mattn/go-sqlite3"
  "time"
)

var(
  G_engine *xorm.Engine
)

func InitDatabase() (err error) {
  var(
    userId int
  )

  //连接
  if G_engine,err = xorm.NewEngine(config.G_config.DbType, config.G_config.DbSource); err != nil {
    return err
  }

  G_engine.DatabaseTZ = time.Local // 必须
  G_engine.TZLocation = time.Local // 必须

  //打印sql日志
  G_engine.ShowSQL(true)

  //同步表
  if err = sycTable();err != nil{
    return
  }
  //初始化数据
  if userId,err = initAdminUser();err != nil {
    return
  }else{
    if _,err = initAdminSecret(userId);err != nil {
      return
    }
  }

  return
}

func sycTable() (err error) {

  if err = G_engine.Sync2(
    new(models.OauthClient),
    new(models.OauthAuthorizationCode),
    new(models.OauthAccessToken),
    new(models.OauthRefreshToken),
    new(models.OauthUser));err != nil {
  }
  return
}

func initAdminUser() (userId int,err error) {
  var(
    userM models.OauthUser
    nUserM  models.OauthUser
    has bool
  )
  userM = models.OauthUser{
    Email: config.G_config.AdminEmail,
  }
  if has,err = G_engine.Get(&userM);err != nil {
    return
  }

  //是否已经创建了管理员
  if !has {
    //创建管理员
    nUserM = models.OauthUser{
      Email: config.G_config.AdminEmail,
      Password: utils.CreatePassword(config.G_config.AdminPassword),
      OpenId: config.G_config.AdminOpenId,
      Name: "管理员",
    }
    if _,err = G_engine.InsertOne(&nUserM);err != nil {
      return
    }

    userId = nUserM.BaseModel.Id
    return
  }else{
    //更新管理员
    nUserM = models.OauthUser{
      Email: config.G_config.AdminEmail,
      Password: utils.CreatePassword(config.G_config.AdminPassword),
      OpenId: config.G_config.AdminOpenId,
    }
    if _,err = G_engine.Where("email=?",nUserM.Email).Update(&nUserM);err != nil {
      return
    }

    userId = userM.BaseModel.Id
    return
  }

}

func initAdminSecret(userId int) (clientId int, err error) {
  var(
    clientM models.OauthClient
    nClientM  models.OauthClient
    has bool
  )
  clientM = models.OauthClient{
    UserId: userId,
  }
  if has,err = G_engine.Get(&clientM);err != nil {
    return
  }
  if !has {
    nClientM = models.OauthClient{
      UserId: userId,
      ClientKey: config.G_config.AdminClientKey,
      ClientSecret:config.G_config.AdminClientSecret,
      RedirectUrl: config.G_config.Localhost,
      AppName: config.G_config.AppName,
    }
    if _,err = G_engine.InsertOne(&nClientM);err != nil {
      return
    }
  }else{
    nClientM = models.OauthClient{
      UserId: userId,
      ClientKey: config.G_config.AdminClientKey,
      ClientSecret:config.G_config.AdminClientSecret,
      RedirectUrl: config.G_config.Localhost,
      AppName: config.G_config.AppName,
    }
    if _,err = G_engine.Where("user_id=?",userId).Update(&nClientM);err != nil {
      return
    }
  }
  clientId = clientM.BaseModel.Id
  return
}
