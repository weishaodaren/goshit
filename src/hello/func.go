package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

func main() {
	params := map[string]interface{}{
		"name": "Tom",
		"pwd":  "123456",
		"age":  30,
	}
	fmt.Printf("sign: %s\n", createSign(params))

}

// MD5
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 签名
func createSign(params map[string]interface{}) string {
	var key []string
	var str string

	for k := range params {
		key = append(key, k)
	}
	fmt.Println("默认的key:", key)
	sort.Strings(key)
	fmt.Println("排序后的key:",key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
			fmt.Println(str)
		} else {
			str = str + fmt.Sprintf("&xl_%v=%v", key[i], params[key[i]])
			fmt.Println("!=0",str)
		}
	}
	// 自定义密钥
	secret := "123456789"

	// 自定义签名算法
	sign := MD5(MD5(str) + MD5(secret))

	return sign
}
