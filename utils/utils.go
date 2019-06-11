package utils

import (
  "math/rand"
  "time"
  "github.com/liu578101804/kun_auth2.0/env"
)

func CreateUserOpenId() string {
  return GetRandomString(8)
}

func CreatePassword(password string) (res, salt string)  {
  salt = Sha1(time.Now().Format(env.TIME_LAYOUT)+GetRandomString(8))
  res = Sha256(password+salt)
  return
}

func GetPassword(password,salt string) string {
  return Sha256(password+salt)
}


func CreateCode() string {
  return Sha1(time.Now().Format(env.TIME_LAYOUT)+GetRandomString(18))[:20]
}

func CreateToken() string {
  return Sha1(time.Now().Format(env.TIME_LAYOUT)+GetRandomString(20))[:40]
}

func CreateClientKey() string {
  return Sha1(time.Now().Format(env.TIME_LAYOUT)+GetRandomString(10))[:20]
}

func CreateClientSecret() string {
  return Sha256(time.Now().Format(env.TIME_LAYOUT)+GetRandomString(10))[:40]
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
  return time.ParseInLocation(env.TIME_LAYOUT ,timeStr,time.Local)
}
