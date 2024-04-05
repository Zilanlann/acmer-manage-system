package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zilanlann/acmer-manage-system/server/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSetting.Secret)

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenTokens(userID int, username string, role string) (aToken, rToken string, err error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(setting.JwtSetting.LongExpiresTime).Unix(), // 过期时间
			Issuer:    setting.JwtSetting.Issuer,                                 // 签发人
		},
	}
	fmt.Printf("claims.ExpiresAt: %v\n", claims.ExpiresAt)
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成 aToken
	aToken, err = token.SignedString(jwtSecret)
	if err != nil {
		log.Println(err)
	}

	// rToken 不需要存储任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(setting.JwtSetting.ShortExpiresTime).Unix(), // 过期时间
		Issuer:    setting.JwtSetting.Issuer,                                  // 签发人
	}).SignedString(jwtSecret)
	if err != nil {
		log.Println(err)
	}

	return aToken, rToken, nil
}

// ParseToken parsing token
func ParseToken(access_token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(access_token, &Claims{}, KeyFunc)

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 第一步 : 判断 rToken 格式对的，没有过期的
// 第二步 : 判断 aToken 格式对的，但是是过期的
// 第三步 : 生成双 token
func RefreshToken(aToken, rToken string) (newToken, newrToken string, err error) {
	// 第一步 : 判断 rToken 格式对的，没有过期的
	if _, err := jwt.Parse(rToken, KeyFunc); err != nil {
		return "", "", err
	}

	// 第二步：从旧的 aToken 中解析出 cliams 数据
	var claims Claims
	_, err = jwt.ParseWithClaims(aToken, &claims, KeyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当 access token 是过期错误，并且 refresh token 没有过期就创建一个新的 access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenTokens(claims.UserID, claims.Username, claims.Role)
	}
	return "", "", err
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	return jwtSecret, nil
}

func Test() {
	aToken, rToken, err := GenTokens(1, "zilanlann", "admin")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aToken)
	fmt.Println(rToken)
	claim, _ := ParseToken(aToken)
	fmt.Println(claim)
}
