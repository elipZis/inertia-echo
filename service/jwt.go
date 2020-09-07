package service

import (
	"elipzis.com/inertia-echo/repository/model"
	"elipzis.com/inertia-echo/util"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"strconv"
	"time"
)

// custom claims extending default ones
type JWTCustomClaims struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	jwt.StandardClaims
}

var JWTSecret = []byte(util.GetEnvOrDefault("JWT_SECRET", "supersecretjwtsecret"))
var JWTLifetime time.Duration

//
func init() {
	lifetime, err := strconv.Atoi(util.GetEnvOrDefault("JWT_LIFETIME_HOURS", "24"))
	if err != nil {
		lifetime = 720
	}
	JWTLifetime = time.Hour * time.Duration(lifetime)
}

//
func (this *Service) GenerateToken(user *model.User) *string {
	claims := &JWTCustomClaims{
		user.Id,
		user.Email,
		user.FirstName,
		user.LastName,
		jwt.StandardClaims{
			Issuer:    util.GetEnvOrDefault("NAME", "elipzis.com/inertia-echo"),
			ExpiresAt: time.Now().Add(JWTLifetime).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and return it
	t, _ := token.SignedString(JWTSecret)
	return &t
}
