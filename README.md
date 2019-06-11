# kun_auth2.0

基于Go语言，实现auth2.0

## 安装

glide安装方法

```
go get github.com/Masterminds/glide
go install github.com/Masterminds/glide
```

安装依赖：

```
glide install
```



## 界面

### 1、注册用户

- `/oauth/register`


### 2、登录用户

- `/oauth/authorize`

## API

#### 根据code获取token

url: `/oauth/access_token`

请求方式： post

请求参数：
- code
- client_secret
- clientId


#### 添加授权信息

url: `/client/add`

请求方式： post

请求参数：
- token
- app_name
- redirect_url


#### token验证

url: `/oauth/check_token`

请求方式： post

请求参数：
- token


#### 获取用户信息

url: `/user`

请求方式： post

请求参数：
- token

