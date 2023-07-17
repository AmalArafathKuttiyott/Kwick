package jwt

import (
	"kwick/helper/database"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
)

func ParseJwt(t string) bool {
	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(t, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return false

	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		_, exists := database.GetUserById(claims.Id)
		if !exists {
			return false
		}
	}
	return true
}
func ExtractBearerToken(t string) string {
	// Check if the Authorization header is present
	if t == "" {
		return ""
	}
	// Split the Authorization header value
	// The token is expected to be in the format: Bearer <token>
	authValue := "Bearer "
	token := t[len(authValue):]
	return token
}
func GeneratJwt(uid uint) (string, bool) {
	secretKey := os.Getenv("SECRET_KEY")
	id := strconv.FormatUint(uint64(uid), 10)
	// Create the Claims
	claims := &jwt.StandardClaims{
		Id: id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", false
	}
	return ss, true
}
func AdminParseJwt(t string) bool {
	tokenString := t

	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return false

	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		user, exists := database.GetUserById(claims.Id)
		if !exists {
			return false
		}
		if !user.IsAdmin {
			return false
		}

	}
	return true
}
func GetUserFromJwt(t string) string {
	tokenString := strings.Split(t, " ")
	tok := tokenString[1]
	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tok, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return ""

	}
	var id string
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		_, exists := database.GetUserById(claims.Id)
		if !exists {
			return ""
		}
		id = claims.Id
	}
	return id
}
