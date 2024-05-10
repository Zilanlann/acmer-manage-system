package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/setting"
)

var jwtSecret = []byte(setting.JwtSetting.Secret)

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenTokens(userID int, username string, role string) (aToken, rToken, exTime string, err error) {
	expireTime := time.Now().Add(setting.JwtSetting.ShortExpiresTime)
	exTime = expireTime.Format("2006-01-02 15:04:05")
	claims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),         // 过期时间
			Issuer:    setting.JwtSetting.Issuer, // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成 aToken
	aToken, err = token.SignedString(jwtSecret)
	if err != nil {
		global.LOG.Error(err.Error())
	}

	// 生成 rToken
	claims.StandardClaims.ExpiresAt = time.Now().Add(setting.JwtSetting.LongExpiresTime).Unix()
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		global.LOG.Error(err.Error())
	}
	return aToken, rToken, exTime, nil
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

func RefreshToken(rToken string) (newAToken, newRToken, exTime string, err error) {
	var claims Claims
	_, err = jwt.ParseWithClaims(rToken, &claims, KeyFunc)
	if err != nil {
		return "", "", "", err
	} else {
		return GenTokens(claims.UserID, claims.Username, claims.Role)
	}
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	return jwtSecret, nil
}
