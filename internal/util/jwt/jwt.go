package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/constraints"
)

func Sign[Integer constraints.Integer](key string, userId Integer) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	return token.SignedString([]byte(key))
}

// Verify verifies the token and returns the claims if the token is valid
// otherwise returns nil and error
func Verify(token, key string) (any, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		return claims, nil
	}
	return nil, err
}

// ExtractPayload extracts the payload from the jwt token
// Note:
// It will not verify the token
// Please use Verify() before calling this function
func ExtractPayload(reqToken string) (map[string]any, error) {
	token := strings.Split(reqToken, "Bearer ")[1]
	rawPayload := strings.Split(token, ".")[1]
	decodedBytes, err := base64.RawURLEncoding.DecodeString(rawPayload)
	if err != nil {
		return nil, err
	}
	var payload map[string]any
	err = json.Unmarshal(decodedBytes, &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
