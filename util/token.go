package util

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Claims struct {
	ID uint
	Username string
	jwt.StandardClaims
}

//生成token
func CreateToken(id uint,username string) (tokenstr string,err error) {

	if err != nil {
		return "", err
	}
	c := &Claims{
		ID: id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "lxy",
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenstr, _ = token.SignedString([]byte(os.Getenv("APPSECRET")))
	return tokenstr,nil

}

//解析token
func ParseToken(tokenstr string) (id uint,username string,err error) {

	c := &Claims{}

	_, err = jwt.ParseWithClaims(tokenstr, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("APPSecret")), nil
	})
	if err != nil {
		return 0,"", err
	}
	return c.ID, c.Username,err
}

