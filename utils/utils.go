package utils

import (
  "math/rand"
  "time"
  "strings"
)

func CreateCode() string {
  return GetRandomString(18)
}

func CreateToken() string {
  return GetRandomString(18)
}

func CreateUserOpenId() string {
  return GetRandomString(18)
}

func CreateClientSecret() string {
  return strings.ToLower(GetRandomString(20))
}

func CreateClientKey() string {
  return strings.ToLower(GetRandomString(40))
}

func CreatePassword(password string) string  {
  return password
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

func GetTime() time.Time {
  var cstZone = time.FixedZone("CST", 8*3600)       // 东八
  return time.Now().In(cstZone)
}

func Time(timeStr string) (time.Time, error) {
  return time.ParseInLocation("2006-01-02 15:04:05" ,timeStr,time.Local)
}
