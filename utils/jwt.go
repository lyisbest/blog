package utils

import (
	"blog/configuration"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"time"
)

type customClaim struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

var (
	secret                  = []byte("iwillrich")
	expTime   time.Duration = 30 * time.Minute
	tokenUsed string        = "1"
)

func GenerateToken(name string) (string, error) {
	claim := customClaim{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(expTime.Seconds()),
			Issuer:    "man",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(secret)
}

func ParseToken(token string) (*jwt.Token, *customClaim, error) {
	claim := &customClaim{}
	t, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	fmt.Printf("parse the token: %v\n", t)
	return t, claim, err
}

func ValidateToken(token string) bool {
	_, _, err := ParseToken(token)
	if err != nil {
		if tp, ok := err.(*jwt.ValidationError); ok {
			if tp.Errors&jwt.ValidationErrorExpired != 0 {
				return false
			}
		}
	}
	return true
}

func RefreshToken(token string) (string, error) {
	_, claim, err := ParseToken(token)
	if err != nil {
		return "", err
	}
	newToken, err := GenerateToken(claim.Name)
	if err != nil {
		return "", err
	}
	return newToken, nil
}

func MarkTokenUsed(token string) error {
	err := configuration.RedisClient.Set(token, tokenUsed, expTime).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetTokenUsed(token string) error {
	result, err := configuration.RedisClient.Get(token).Result()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return err
	}
	if result == tokenUsed {
		return BlogError{ErrorCode: -1008, ErrorMessage: "token has security risks. Please change your password in time."}
	}
	return nil
}
