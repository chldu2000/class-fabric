package cryptox

import (
	"medicineApp/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims jwt 中的其他信息
type Claims struct {
	jwt.StandardClaims

	Name string `json:"name"`
}

// ClaimsAdmin jwt 中的其他信息
type ClaimsAdmin struct {
	jwt.StandardClaims

	SystemAccount string `json:"system_account"`
}

// GenerateToken 生成 jwt
func GenerateToken(account string) (string, error) {
	cfg := config.C.JWT

	jwtSecret := []byte(cfg.Secret)
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(cfg.Expires * int(time.Hour)))

	claims := Claims{
		Name: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 格式化 JWT
func ParseToken(token string) (*Claims, error) {
	cfg := config.C.JWT

	jwtSecret := []byte(cfg.Secret)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, err
		}
	}

	return nil, err
}
