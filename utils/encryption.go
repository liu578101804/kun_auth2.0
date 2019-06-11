package utils

import (
  "crypto/sha1"
  "crypto/sha256"
  "fmt"
  "crypto/md5"
)

func Sha1(input string) string {
  hash := sha1.New()
  hash.Write([]byte(input))
  result := hash.Sum(nil)
  return fmt.Sprintf("%x",result)
}

func Sha256(input string) string {
  hash := sha256.New()
  hash.Write([]byte(input))
  result := hash.Sum(nil)
  return fmt.Sprintf("%x",result)
}

func Md5(input string) string {
  hash := md5.New()
  hash.Write([]byte(input))
  result := hash.Sum(nil)
  return fmt.Sprintf("%x",result)
}
