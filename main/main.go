package main

import (
  "github.com/liu578101804/kun_auth2.0/config"
  "flag"
  "runtime"
  "fmt"
  "github.com/liu578101804/kun_auth2.0/database"
)

var (
  confFile string //配置文件路径
)

// 解析命令行参数
func initArgs()  {
  flag.StringVar(&confFile, "config","./conf.json","指定系统的配置文件")
  flag.Parse()
}

// 初始化线程数量
func initEnv() {
  runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
  var(
    err error
  )

  //初始化命令行参数
  initArgs()

  //初始化线程
  initEnv()

  //加载配置
  if err = config.InitConfig(confFile);err != nil {
    goto ERR
  }


  //初始化数据库
  if err = database.InitDatabase();err != nil {
    goto ERR
  }


  return
ERR:
  fmt.Println(err)
}
