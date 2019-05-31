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
