package helpers

import (
	"errors"
	"os"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID string `json:"id"`
	Email  string `json:"email"`

	jwt.RegisteredClaims
}

var jwtKey []byte

func SetJWTKey(key string) {
	jwtKey = []byte(key)
}

func GetJWTKey() []byte {
	return []byte(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	secretKey := GetJWTKey()

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenerateTokens(email, userID string) (string, string, error) {

	claims := &Claims{
		Email:  email,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	refreshClaims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedAccesToken, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return signedAccesToken, signedRefreshToken, nil

}

func HashPassword(password *string) *string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	hashedPassword := string(bytes)
	return &hashedPassword
}

func VerifyPassword(foundPwd, pwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(foundPwd), []byte(pwd))

	return err == nil, err
}

func CheckPasswordStrength(pwd string) error {
	// abcUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// abcLower := strings.ToLower(abcUpper)
	// numericChars := ""
	//
	if len(pwd) < 8 {
		return errors.New("password length should be at least 8 characters long")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(pwd) {
		return errors.New("password length should contain at least an uppercase letter")
	}

	// At least one lowercase letter
	if !regexp.MustCompile(`[a-z]`).MatchString(pwd) {
		return errors.New("password length should contain at least a lowercase letter")
	}

	// At least one digit
	if !regexp.MustCompile(`[0-9]`).MatchString(pwd) {
		return errors.New("password length should contain at least a numeric character")
	}

	// At least one special character (adjust as needed)
	if !regexp.MustCompile(`[!@#$%^&*()]`).MatchString(pwd) {
		return errors.New("password length should contain at least a symbol !@#$%^&*()")
	}

	return nil
}

// get environment go app is running
func IsProd() bool {
	prod := os.Getenv("APP_ENV")

	return prod == "production"
}
