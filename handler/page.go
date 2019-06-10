package handler

import (
  "github.com/gin-gonic/gin"
  "github.com/liu578101804/kun_auth2.0/config"
  "net/http"
  "fmt"
  "github.com/liu578101804/kun_auth2.0/service"
  "github.com/liu578101804/kun_auth2.0/models"
  "errors"
  "github.com/liu578101804/kun_auth2.0/utils"
)

func RegPageHandler()  {

  //页面
  //登录
  G_app.GET("/", HomePage)
  G_app.GET("/oauth/login", LoginPage)
  //注册
  G_app.GET("/oauth/register", RegisterPage)

  //操作
  //登录
  G_app.POST("/oauth/login", Login)
  //注册
  G_app.POST("/oauth/register", Register)
  //通过code换token
  G_app.POST("/oauth/token", OauthToken)
  ////刷新token
  //G_app.POST("/oauth/token/ref", _)
  ////验证token
  //G_app.POST("/oauth/token/check", _)
  ////获取用户信息
  //G_app.POST("/user/info", _)

}

func HomePage(c *gin.Context)  {
  c.HTML(http.StatusOK,"home/index", gin.H{
    "title":"首页",
  })
}

func LoginPage(c *gin.Context) {
  var(
   clientId  string
   //returnTo  string
  )
  clientId = c.DefaultQuery("client_id","")
  //returnTo = c.DefaultQuery("return_to","")

  if clientId == "" {
    goto ERR
  }

  c.HTML(http.StatusOK, "home/login", gin.H{
    "title":"登录",
  })
  return

  ERR:
    //重定向到管理员
    c.Redirect(http.StatusFound,"/oauth/login?client_id="+config.G_config.AdminClientKey)
}

func RegisterPage(c *gin.Context) {
  c.HTML(http.StatusOK, "home/register", gin.H{
    "title":"注册",
  })
}

func Login(c *gin.Context) {

  var(
    errCode int
    email  string
    password  string
    clientId string
    err error
    userModel *models.OauthUser
    authClient *models.OauthClients
    code string //返回的code

    redirectUrl string
    returnTo string //回调回去时带的参数
  )

  //验证client_id
  clientId = c.DefaultPostForm("client_id","")
  if clientId=="" {
    errCode = 100001
    err = errors.New("client_id不能为空")
    goto ERR
  }
  if authClient,err=service.QueryClientByClientKey(clientId);err != nil {
    errCode = 100002
    goto ERR
  }

  //验证用户名和密码
  email = c.DefaultPostForm("email","")
  password = c.DefaultPostForm("password","")
  if email==""||password=="" {
    errCode = 100003
    err = errors.New("email,password不能为空")
    goto ERR
  }

  //登录操作
  if userModel,err = service.LoginService(email, password); err != nil {
    errCode = 100004
    goto ERR
  }

  //生成code
  code = utils.CreateCode()
  if err = service.CreateCode(userModel.BaseModel.Id, authClient.ClientKey, code); err != nil {
    errCode = 100005
    goto ERR
  }

  redirectUrl = authClient.RedirectUrl + "?code="+ code
  // 获取可选参数
  returnTo = c.DefaultPostForm("return_to","")
  redirectUrl = redirectUrl+"&return_to=" + returnTo

  c.JSON(http.StatusOK, createSuccess(gin.H{
    "redirect_url": redirectUrl,
  }))

  return
  ERR:
    c.JSON(http.StatusOK, createERR(errCode,err.Error()))
}

func Register(c *gin.Context)  {

  var(
    email  string
    password  string
    userModel *models.OauthUser
    err error
    errCode int
  )
  email = c.DefaultPostForm("email","")
  password = c.DefaultPostForm("password","")

  if userModel,err = service.RegisterService(email,password);err != nil {
    errCode = 300001
    goto ERR
  }

  c.JSON(http.StatusOK, createSuccess(gin.H{
    "id": userModel.BaseModel.Id,
    "open_id": userModel.OpenId,
  }))

  return
  ERR:
    c.JSON(http.StatusOK, createERR(errCode,err.Error()))
}

func OauthToken(c *gin.Context) {

  var(
    errCode int
    clientSecret  string
    code  string
    err error
    authorization *models.OauthAuthorizationCode
    client *models.OauthClients
    accessToken *models.OauthAccessToken
  )

  clientSecret = c.DefaultPostForm("client_secret","")
  code = c.DefaultPostForm("code","")

  if clientSecret==""||code=="" {
    errCode = 200001
    err = errors.New("client_secret,code不能为空")
    goto ERR
  }

  //查询code信息
  if authorization,err = service.GetAuthorizationCodeByCode(code); err != nil {
    errCode = 200002
    goto ERR
  }
  //查关于client的信息
  if client,err = service.QueryClientByClientKey(authorization.ClientId); err != nil{
    errCode = 200003
    goto ERR
  }
  if client.ClientSecret != clientSecret {
    errCode = 200004
    err = errors.New("秘钥不对")
    goto ERR
  }

  //生成token
  if accessToken,err = service.CreateAccessToken(client.ClientKey,authorization.UserId);err != nil {
    fmt.Println(err)
    return
  }

  c.JSON(http.StatusOK, createSuccess(gin.H{
    "token": accessToken.Token,
    "expires_at": accessToken.ExpiresAt.Format("2006-01-02 15:04:05"),
  }))

  return
  ERR:
    c.JSON(http.StatusOK, createERR(errCode,err.Error()))
}
