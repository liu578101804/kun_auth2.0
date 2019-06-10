package handler

import (
  "github.com/gin-gonic/gin"
)

func createERR(code int,msg string) *gin.H {
  return &gin.H{
    "code": code,
    "msg": msg,
    "data":"",
  }
}

func createSuccess(data interface{}) *gin.H {
  return &gin.H{
    "code": 200,
    "msg": "操作成功",
    "data": data,
  }
}
