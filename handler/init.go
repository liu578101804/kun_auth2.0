package handler

import (
  "fmt"
  "github.com/liu578101804/kun_auth2.0/config"
  "github.com/gin-gonic/gin"
  "net/http"
)

var(
  G_app *gin.Engine
)

func Init() (err error) {

  //var(
  //  logF *os.File
  //)
  //
  ////配置日志
  //if logF,err = os.Create(config.G_config.LogoPath);err != nil {
  //  return err
  //}
  ////配置日志
  //gin.DefaultWriter = io.Writer(logF)

  //创建
  G_app = gin.Default()

  //自定义分隔符
  G_app.Delims("{{", "}}")

  //设置静态资源路径
  G_app.Static("/static", config.G_config.StaticPath)
  G_app.StaticFile("/favicon.ico", config.G_config.StaticPath+"/favicon.ico")
  //配置全局模板路径
  G_app.LoadHTMLGlob(config.G_config.HtmlPath+"/*/**")

  //注册页面路由
  RegHomeHandler()
  RegAdminHandler()
  RegOauthHandler()
  RegUserHandler()

  //启动运行
  fmt.Println("server is run listing at", config.G_config.ApiPort)
  err = G_app.Run(fmt.Sprintf(":%d", config.G_config.ApiPort))

  return
}

func RegHomeHandler() {
  G_app.GET("/", HomePage)
}

func HomePage(c *gin.Context) {
  c.HTML(http.StatusOK,"home/index", gin.H{
    "title":"首页",
  })
}



