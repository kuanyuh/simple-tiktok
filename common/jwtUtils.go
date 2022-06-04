package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kuanyuh/simple-tiktok/service"
)

const (
	// SIGNED_KEY HS256 signed key
	SIGNED_KEY = "A40CCD9tC3553oer3dj26588E"
)

//GetHStoken 获取签名算法为HS256的token
func GetHStoken(tokenFirst string, user *service.User) string {
	claims := jwt.MapClaims{ //claims里把userId和username封装成map
		"tokenES": tokenFirst,
		//解析时，该变量的类型被转换成float64
		"id":       user.Id,
		"username": user.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//加密算法是HS256时，这里的SignedString必须是[]byte类型
	ss, err := token.SignedString([]byte(SIGNED_KEY))
	if err != nil {
		fmt.Printf("token生成签名错误,err=%v\n", err)
		return ""
	}
	return ss
}

//ParseHStoken 解析签名算法为HS256的token
func ParseHStoken(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGNED_KEY), nil
	})
	if err != nil {
		fmt.Println("HS256的token解析错误，err:", err)
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("ParseHStoken:claims类型转换失败")
		return nil
	}
	return claims
}
