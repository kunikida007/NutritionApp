package util

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/subosito/gotenv"
)

// カスタムのJWTクレーム
type MyJWTClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

// echo-jwtの設定
var Config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(MyJWTClaims)
	},
	SigningKey: getJWTSecret(),
}

// JWTの署名キーを取得
func getJWTSecret() []byte {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

// 指定されたユーザーIDと名前から署名付きのJWTトークン文字列を生成
func GenerateSignedString(userId int, name string) (string, error) {
	claims := &MyJWTClaims{
		userId,
		name,
		jwt.RegisteredClaims{
			//  有効期限72時間
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

func GetUserIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*MyJWTClaims)
	return claims.ID
}
