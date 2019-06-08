package handler

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
  "github.com/liu578101804/kun_auth2.0/service"
  "github.com/liu578101804/kun_auth2.0/models"
  "errors"
  "github.com/liu578101804/kun_auth2.0/utils"
)


func RegPageHandler()  {

  G_app.GET("/login", LoginLoginPage)
  G_app.POST("/login", Login)
  //通过code换token
  G_app.POST("/oauth/token", OauthToken)

}

func LoginLoginPage(c *gin.Context) {

  var(
    clientId  string
    returnTo  string
  )
  clientId = c.DefaultQuery("client_id","")
  returnTo = c.DefaultQuery("return_to","")

  fmt.Println(clientId,returnTo)

  c.HTML(http.StatusOK, "home/login", gin.H{
    "title":"blog",
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
    goto RJSON
  }
  if authClient,err=service.QueryClientByClientKey(clientId);err != nil {
    errCode = 100002
    goto RJSON
  }

  //验证用户名和密码
  email = c.DefaultPostForm("email","")
  password = c.DefaultPostForm("password","")
  if email==""||password=="" {
    errCode = 100003
    err = errors.New("email,password不能为空")
    goto RJSON
  }

  //登录操作
  if userModel,err = service.LoginService(email, password); err != nil {
    errCode = 100004
    goto RJSON
  }

  //生成code
  code = utils.CreateCode()
  if err = service.CreateCode(userModel.BaseModel.Id, authClient.ClientKey, code); err != nil {
    errCode = 100005
    goto RJSON
  }


  redirectUrl = authClient.RedirectUrl + "?code="+ code
  // 获取可选参数
  returnTo = c.DefaultPostForm("return_to","")
  redirectUrl = redirectUrl+"&return_to=" + returnTo

  c.JSON(http.StatusOK, gin.H{
    "code": "200",
    "msg": "操作成功",
    "data": gin.H{
      "redirect_url": redirectUrl,
    },
  })

  return

RJSON:

  c.JSON(http.StatusOK, gin.H{
    "code": errCode,
    "msg": err.Error(),
    "data": "",
  })

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
    goto RJSON
  }

  //查询code信息
  if authorization,err = service.GetAuthorizationCodeByCode(code); err != nil {
    errCode = 200002
    goto RJSON
  }
  //查关于client的信息
  if client,err = service.QueryClientByClientKey(authorization.ClientId); err != nil{
    errCode = 200003
    goto RJSON
  }
  if client.ClientSecret != clientSecret {
    errCode = 200004
    err = errors.New("秘钥不对")
    goto RJSON
  }

  //生成token
  if accessToken,err = service.CreateAccessToken(client.ClientKey,authorization.UserId);err != nil {
    fmt.Println(err)
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "code": 200,
    "msg": "操作成功",
    "data": gin.H{
      "token": accessToken.Token,
      "expires_at": accessToken.ExpiresAt.Format("2006-01-02 15:04:05"),
    },
  })

  return

RJSON:

  c.JSON(http.StatusOK, gin.H{
    "code": errCode,
    "msg": err.Error(),
    "data": "",
  })
}
