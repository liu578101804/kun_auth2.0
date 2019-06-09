package config

import (
  "io/ioutil"
  "encoding/json"
)

// 程序配置
type Config struct {
  ApiPort     int       `json:"api_port"`
  DbType      string    `json:"db_type"`
  DbSource    string    `json:"db_source"`
  StaticPath  string    `json:"static_path"`
  HtmlPath    string    `json:"html_path"`
  LogoPath    string    `json:"logo_path"`

  TokenExpiresAt  int   `json:"token_expires_at"`
  CodeExpiresAt   int   `json:"code_expires_at"`

  AppName   string    `json:"app_name"`
  Localhost     string  `json:"localhost"`
  LoadLocation  string  `json:"load_location"`

  AdminClientKey  string  `json:"admin_client_key"`
  AdminClientSecret string  `json:"admin_client_secret"`
  AdminEmail  string  `json:"admin_email"`
  AdminPassword   string  `json:"admin_password"`
  AdminOpenId     string  `json:"admin_open_id"`
}

var (
  G_config *Config
)

func InitConfig(filename string) (err error) {

  var (
    content []byte
    conf Config
  )

  // 1, 把配置文件读进来
  if content, err = ioutil.ReadFile(filename); err != nil {
    return
  }

  // 2, 做JSON反序列化
  if err = json.Unmarshal(content, &conf); err != nil {
    return
  }

  // 3, 赋值单例
  G_config = &conf

  return
}
