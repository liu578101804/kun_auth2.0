package models

import "time"

type ModelCommon struct {
  Id              int         `xorm:"not null pk autoincr" json:"id" `
  CreatedAt       time.Time   `xorm:"created" json:"created_at"`
  UpdatedAt       time.Time   `xorm:"updated " json:"updated_at"`
  DeletedAt       time.Time   `xorm:"deleted" json:"deleted_at"` //comment(主键)
}

type OauthClients struct {
  TempUser     ModelCommon  `xorm:"extends"`
  ClientKey    string       `xorm:"VARCHAR(255)"    json:"client_key"`
  ClientSecret string       `xorm:"VARCHAR(255)"    json:"client_secret"`
  RedirectUrl  string       `xorm:"VARCHAR(200)"    json:"redirect_uri"`
}

func (c *OauthClients) TableName() string {
  return "oauth_clients"
}

