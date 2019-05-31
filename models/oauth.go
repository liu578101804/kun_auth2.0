package models

import "time"

type ModelCommon struct {
  Id              int         `xorm:"not null pk INT(10) autoincr" json:"id"`
  CreatedAt       time.Time   `json:"created_at"`
  UpdatedAt       time.Time   `json:"updated_at"`
  DeletedAt       time.Time   `json:"deleted_at"`
}

type OauthClients struct {
  ModelCommon  `xorm:"extends"`
  ClientKey    string       `xorm:"varchar(255) client_key"       json:"client_key"`
  ClientSecret string       `xorm:"varchar(255) client_secret"    json:"client_secret"`
  RedirectURI  string       `xorm:"varchar(200) redirect_uri"     json:"redirect_uri"`
}

func (c *OauthClients) TableName() string {
  return "oauth_clients"
}
