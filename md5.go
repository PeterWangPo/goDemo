package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Common_md5(md5_string string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(md5_string))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func main() {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("test md5 encrypto"))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Print(hex.EncodeToString(cipherStr))
}
