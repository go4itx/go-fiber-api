package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"home/pkg/utils/conf"
	"home/pkg/utils/xcast"
	"log"
	"time"
)

var Config conf.Jwt

func init() {
	if err := conf.Load("auth.jwt", &Config); err != nil {
		log.Printf("jwt init error: %v \n", err)
	}
}

type User struct {
	ID     uint   `json:"id"`
	RoleID uint   `json:"roleID"`
	Name   string `json:"name"`
}

// CreateToken return token
func CreateToken(user User) (token string, err error) {
	claims := jwt.MapClaims{
		"id":   user.ID,
		"aud":  user.Name,
		"iss":  Config.Iss,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * time.Duration(24*Config.Exp)).Unix(),
		"role": user.RoleID,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(Config.Signing))
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
		ID:     uint(xcast.ToInt64(claims["id"])),
		RoleID: uint(xcast.ToInt64(claims["role"])),
		Name:   xcast.ToString(claims["aud"]),
	}

	return
}
