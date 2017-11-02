package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	fmt.Println("vim-go")
}

// hard code secret
var Secret string = "I am a hero"

func NewToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected sigining method: %v", token.Header["alg"])
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return "", false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), true
	} else {
		return "", false
	}
}
