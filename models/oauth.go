package models

import "time"

type ModelCommon struct {
  Id              int         `xorm:"not null pk autoincr" json:"id" `
  CreatedAt       time.Time   `xorm:"created" json:"created_at"`
  UpdatedAt       time.Time   `xorm:"updated " json:"updated_at"`
  DeletedAt       time.Time   `xorm:"deleted" json:"deleted_at"` //comment(主键)
}

// 机构
type OauthClient struct {
  BaseModel    ModelCommon  `xorm:"extends"`
  UserId        int         `xorm:"INT(10) 'user_id'" json:"user_id"` //创建人
  ClientKey    string       `xorm:"VARCHAR(255)"    json:"client_key"`
  ClientSecret string       `xorm:"VARCHAR(255)"    json:"client_secret"`
  RedirectUrl  string       `xorm:"VARCHAR(200)"    json:"redirect_uri"`  //回调地址
  AppName      string      `xorm:"VARCHAR(50)"     json:"app_name"`  //应用名字
}

func (c *OauthClient) TableName() string {
  return "oauth_client"
}

// 用户
type OauthUser struct {
  BaseModel      ModelCommon  `xorm:"extends"`
  Email         string    `xorm:"VARCHAR(50)" json:"email"`
  Password      string    `xorm:"VARCHAR(64)" json:"-"`
  Salt          string    `xorm:"VARCHAR(40)" json:"-"`

  OpenId        string    `xorm:"VARCHAR(32) unique 'open_id'" json:"open_id"`
  Name          string    `xorm:"VARCHAR(50)" json:"name"`
  Phone         string    `xorm:"VARCHAR(20)" json:"phone"`
}

func (c *OauthUser) TableName() string {
  return "oauth_user"
}

// 登录后获得的code
type OauthAuthorizationCode struct {
  BaseModel     ModelCommon  `xorm:"extends"`
  ClientId     string    `xorm:"VARCHAR(60) 'client_id'" json:"client_id"`
  UserId       int       `xorm:"INT(10)" json:"user_id"`
  Code         string    `xorm:"VARCHAR(60)" json:"code"`
  ExpiresAt     time.Time    `json:"expires_at"`
  IsUse        string     `xorm:"CHAR(1)" json:"is_use"`
}

func (c *OauthAuthorizationCode) TableName() string {
  return "oauth_authorization_code"
}

// 用户Token
type OauthAccessToken struct {
  BaseModel     ModelCommon  `xorm:"extends"`
  ClientId      string    `xorm:"VARCHAR(60)" json:"client_id"`
  UserId        int       `json:"user_id"`
  Token         string    `json:"token"`
  ExpiresAt     time.Time    `json:"expires_at"`
}

func (c *OauthAccessToken) TableName() string {
  return "oauth_access_token"
}

// 用户刷新延时Token
type OauthRefreshToken struct {
  BaseModel     ModelCommon  `xorm:"extends"`
  ClientId     string    `xorm:"VARCHAR(60)" json:"client_id"`
  UserId        int       `json:"user_id"`
  Token         string    `json:"token"`
  ExpiresAt     string    `json:"expires_at"`
}

func (c *OauthRefreshToken) TableName() string {
  return "oauth_refresh_token"
}
