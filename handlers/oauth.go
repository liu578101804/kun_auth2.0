package handlers

import (
  "github.com/gin-gonic/gin"
  "github.com/liu578101804/kun_auth2.0/config"
  "net/http"
  "github.com/liu578101804/kun_auth2.0/service"
  "github.com/liu578101804/kun_auth2.0/models"
  "errors"
  "github.com/liu578101804/kun_auth2.0/utils"
  "time"
  "github.com/liu578101804/kun_auth2.0/env"
)

func RegOauthHandler()  {
  var(
    router *gin.RouterGroup
  )
  router = G_app.Group("/oauth")

  //登录
  router.GET("/authorize", LoginPage)
  //注册
  router.GET("/register", RegisterPage)

  //登录
  router.POST("/authorize", Login)
  //注册
  router.POST("/register", Register)
  //通过code换token
  router.POST("/access_token", OauthAccessToken)
  //验证token
  router.POST("/check_token", OauthCheckToken)

}

func LoginPage(c *gin.Context) {
  var(
   clientId  string
  )
  clientId = c.DefaultQuery("client_id","")

  if clientId == "" {
    goto ERR
  }

  c.HTML(http.StatusOK, "home/login", gin.H{
    "title":"登录",
  })
  return

  ERR:
    //重定向到管理员
    c.Redirect(http.StatusFound,"/oauth/authorize?client_id="+config.G_config.AdminClientKey)
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
    authClient *models.OauthClient
    code string //返回的code

    redirectUrl string
    returnTo string //回调回去时带的参数
  )

  //验证client_id
  clientId = c.DefaultPostForm("client_id","")
  if clientId=="" {
    errCode = env.ERRCODE_LOGIN + 1
    err = errors.New("client_id不能为空")
    goto ERR
  }
  if authClient,err=service.QueryClientByClientKey(clientId);err != nil {
    errCode = env.ERRCODE_LOGIN + 2
    goto ERR
  }

  //验证用户名和密码
  email = c.DefaultPostForm("email","")
  password = c.DefaultPostForm("password","")
  if email==""||password=="" {
    errCode = env.ERRCODE_LOGIN + 3
    err = errors.New("email,password不能为空")
    goto ERR
  }

  //登录操作
  if userModel,err = service.LoginService(email, password); err != nil {
    errCode = env.ERRCODE_LOGIN + 4
    goto ERR
  }

  //生成code
  code = utils.CreateCode()
  if err = service.CreateCode(userModel.BaseModel.Id, authClient.ClientKey, code); err != nil {
    errCode = env.ERRCODE_LOGIN + 5
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
    errCode = env.ERRCODE_REGISTER + 1
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

func OauthAccessToken(c *gin.Context) {

  var(
    errCode int
    clientId  string
    clientSecret  string
    code  string
    err error
    authorization *models.OauthAuthorizationCode
    client *models.OauthClient
    accessToken *models.OauthAccessToken
    nowTime time.Time
  )
  nowTime = time.Now()

  clientId = c.DefaultPostForm("client_id","")
  clientSecret = c.DefaultPostForm("client_secret","")
  code = c.DefaultPostForm("code","")

  if clientSecret==""||code=="" {
    errCode = env.ERRCODE_OAUTH_ACCESS_Token + 1
    err = errors.New("client_secret,code不能为空")
    goto ERR
  }

  //查询code信息
  if authorization,err = service.GetAuthorizationCodeByCode(code); err != nil {
    errCode = env.ERRCODE_OAUTH_ACCESS_Token + 2
    goto ERR
  }

  //判断是否过期或者已经使用
  if authorization.IsUse == "T" || authorization.ExpiresAt.Sub(nowTime).Seconds() <= 0 {
    errCode = env.ERRCODE_OAUTH_ACCESS_Token + 3
    err = errors.New("code过期或者以使用")
    goto ERR
  }

  //查关于client的信息
  if client,err = service.QueryClientByClientKey(authorization.ClientId); err != nil{
    errCode = env.ERRCODE_OAUTH_ACCESS_Token + 4
    goto ERR
  }
  if client.ClientSecret != clientSecret || client.ClientKey == clientId {
    errCode = env.ERRCODE_OAUTH_ACCESS_Token + 5
    err = errors.New("秘钥不对")
    goto ERR
  }

  //生成token
  if accessToken,err = service.CreateAccessToken(code, client.ClientKey, authorization.UserId);err != nil {
    errCode = env.ERRCODE_OAUTH_ACCESS_Token + 6
    return
  }

  c.JSON(http.StatusOK, createSuccess(gin.H{
    "token": accessToken.Token,
    "expires_at": accessToken.ExpiresAt.Format(env.TIME_LAYOUT),
  }))

  return
  ERR:
    c.JSON(http.StatusOK, createERR(errCode,err.Error()))
}

// 验证token是否过期
func OauthCheckToken(c *gin.Context)  {
  var(
    token string
    err error
    errCode int
    accessTokenModel *models.OauthAccessToken
  )

  token = c.DefaultPostForm("token","")
  if token == "" {
    errCode = env.ERRCODE_OAUTH_CHECK_TOKEN + 1
    err = errors.New("token参数不能为空")
    goto ERR
  }

  if accessTokenModel, err = service.CheckToken(token);err != nil {
    errCode = env.ERRCODE_OAUTH_CHECK_TOKEN + 2
    goto ERR
  }

  c.JSON(http.StatusOK, createSuccess(gin.H{
    "expires_at": accessTokenModel.ExpiresAt.Format(env.TIME_LAYOUT),
  }))

  return
  ERR:
    c.JSON(http.StatusOK, createERR(errCode,err.Error()))
}
