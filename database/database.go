package database

import (
  _ "github.com/mattn/go-sqlite3"
  "github.com/go-xorm/xorm"

  "github.com/liu578101804/kun_auth2.0/config"
  "github.com/liu578101804/kun_auth2.0/models"
  "github.com/go-xorm/core"
  "fmt"
)

var(
  G_engine *xorm.Engine
)

func InitDatabase() (err error) {

  //连接
  if G_engine,err = xorm.NewEngine(config.G_config.DbType, config.G_config.DbSource); err != nil {
    return err
  }

  //设置表前缀
  //TODO: 好像不生效
  tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "ka_")
  G_engine.SetTableMapper(tbMapper)

  //打印sql日志
  G_engine.ShowSQL(true)

  //同步表
  sycTable()

  return nil
}

func sycTable()  {
  var(
    err error
  )
  if err = G_engine.Sync2(
    new(models.OauthClients));err != nil {
    fmt.Println(err)
  }
}
