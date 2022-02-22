package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"api/src/config"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(ID uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	token, error := jwt.Parse(tokenString, getAuthenticationKey)
	if error != nil {
		return error
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)

	token, error := jwt.Parse(tokenString, getAuthenticationKey)
	if error != nil {
		return 0, error
	}

	permissions, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, error := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if error != nil {
			return 0, error
		}

		return userID, nil
	}

	return 0, error
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func getAuthenticationKey(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("Unexpected signing method")
	}

	return config.SecretKey, nil
}
