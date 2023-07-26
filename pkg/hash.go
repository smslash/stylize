package pkg

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

func CheckHash(hash, file string) bool {
	standard := "ac85e83127e49ec42487f272d9b9db8b"
	shadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	thinkertoy := "db448376863a4b9a6639546de113fa6f"
	if hash == standard && strings.ToLower(file)+".txt" == "standard.txt" {
		return true
	}
	if hash == shadow && strings.ToLower(file)+".txt" == "shadow.txt" {
		return true
	}
	if hash == thinkertoy && strings.ToLower(file)+".txt" == "thinkertoy.txt" {
		return true
	}
	fmt.Println(hash)
	fmt.Println(file)
	return false
}
