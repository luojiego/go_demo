package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var (
	secretKey = "9sh4cvey8mktd0arunj1xz2q"
)

func main() {
	c := &cookieInfo{
		UserName: "root",
		Expire:   time.Now().Add(time.Hour * 24).Unix(),
	}
	tokenString, err := createToken(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenString)

	c1, err := parseToken(tokenString)
	if err != nil {
		panic(err)
	}
	fmt.Println(c1)
}

func createToken(c *cookieInfo) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"user":   c.UserName,
		"expire": strconv.FormatInt(c.Expire, 10),
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func parseToken(tokenString string) (*cookieInfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c := &cookieInfo{}
		if name, ok := claims["user"].(string); ok {
			c.UserName = name
		} else {
			return nil, fmt.Errorf("can't get user name")
		}

		if expire, ok := claims["expire"].(string); ok {
			exp, err := strconv.ParseInt(expire, 10, 64)
			if err != nil {
				return nil, err
			}
			c.Expire = exp
		} else {
			return nil, fmt.Errorf("can't get expire")
		}

		return c, err
	}
	return nil, err
}

type cookieInfo struct {
	UserName string // 用户名
	Expire   int64  // 过期时间
}
