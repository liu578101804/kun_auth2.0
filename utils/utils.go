package utils

import (
  "math/rand"
  "time"
)

func CreateCode() string {
  return GetRandomString(18)
}

func CreateToken() string {
  return GetRandomString(18)
}

// 生成随机数
func GetRandomString(l int) string {
  str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  bytes := []byte(str)
  result := []byte{}
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  for i := 0; i < l; i++ {
    result = append(result, bytes[r.Intn(len(bytes))])
  }
  return string(result)
}
