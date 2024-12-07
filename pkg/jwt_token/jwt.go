package jwttoken

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int64) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	ttlKeyStr := os.Getenv("TTL_KEY")

	ttlKey, err := strconv.Atoi(ttlKeyStr)
	if err != nil {
		return "", fmt.Errorf("Invalid TTL_KEY value: %s", ttlKeyStr)
	}

	expiration := time.Duration(ttlKey) * time.Hour
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(secretKey)
}
