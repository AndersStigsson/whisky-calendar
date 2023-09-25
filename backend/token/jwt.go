package token

import (
	"errors"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AndersStigsson/whisky-calendar/config"
	"github.com/AndersStigsson/whisky-calendar/domain"
	"github.com/golang-jwt/jwt/v5"
)

// JWT Claims struct
type Claims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.RegisteredClaims
}

// Generate new JWT Token
func GenerateJWTToken(user *domain.User) (string, error) {
	// Register the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		ID:       strconv.FormatInt(user.ID, 10),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretKey := config.Get("server.jwtSecretKey")
	// Register the JWT string
	tokenString, err := token.SignedString([]byte(jwtSecretKey.(string)))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Extract JWT From Request
func ExtractJWTFromRequest(r *http.Request) (map[string]interface{}, error) {
	// Get the JWT string
	tokenString := ExtractBearerToken(r)

	// Initialize a new instance of `Claims` (here using Claims map)
	claims := jwt.MapClaims{}

	// Parse the JWT string and repositories the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (jwtKey interface{}, err error) {
		return []byte(config.Get("server.jwtSecretKey").(string)), err
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token ")
	}

	return claims, nil
}

// Extract bearer token from request Authorization header
func ExtractBearerToken(r *http.Request) string {
	headerAuthorization := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerAuthorization, " ")
	if len(bearerToken) > 1 {
		return html.EscapeString(bearerToken[1])
	}
	return ""
}
