package jwt

import (
	conf "home/config"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
)

type Config struct {
	Exp     int64
	Iss     string
	Signing string
}

var (
	config  Config
	Signing string
)

func init() {
	if err := conf.Load("auth.jwt", &config); err != nil {
		log.Printf("jwt init error: %v \n", err)
	}

	Signing = config.Signing
}

type User struct {
	ID     uint   `json:"id"`
	RoleID uint   `json:"roleID"`
	Name   string `json:"name"`
}

// CreateToken return token
func CreateToken(user User) (token string, expire int64, err error) {
	expire = time.Now().Add(time.Hour * time.Duration(24*config.Exp)).Unix()
	claims := jwt.MapClaims{
		"id":   user.ID,
		"aud":  user.Name,
		"iss":  config.Iss,
		"iat":  time.Now().Unix(),
		"exp":  expire,
		"role": user.RoleID,
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Signing))
	return
}

// ParseToken return jwt claims
func ParseToken(jwtToken interface{}) (user User, err error) {
	if jwtToken == nil {
		err = fiber.ErrNetworkAuthenticationRequired
		return
	}

	token := jwtToken.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	user = User{
		ID:     cast.ToUint(claims["id"]),
		RoleID: cast.ToUint(claims["role"]),
		Name:   cast.ToString(claims["aud"]),
	}

	return
}
