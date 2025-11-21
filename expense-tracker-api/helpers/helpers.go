package helpers

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserID string `json:"id"`
	Email  string `json:"email"`

	jwt.StandardClaims
}

var jwtKey []byte

func SetJWTKey(key string) {
	jwtKey = []byte(jwtKey)
}

func GetJWTKey() []byte {
	return []byte(jwtKey)
}

func ValidateToken(tokenString string) (Claims, error) {
	secretKey := GetJWTKey()
}
