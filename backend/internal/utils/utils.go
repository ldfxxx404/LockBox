package utils

import (
	"back/config"
	"errors"
	"flag"
	"time"

	log "github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(userID int, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(config.JWTSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func GetUserID(c *fiber.Ctx) int {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return int(claims["user_id"].(float64))
}

func CheckSize(used, size int64, limit int) error {
	sizeMB := size / (1024 * 1024)
	if used+sizeMB > int64(limit) {
		return errors.New("file size is to large")
	}
	return nil
}

func ParseLoglevelFlags() {
	debug := flag.Bool("debug", false, "Enable debug logging")
	info := flag.Bool("info", false, "Enable info logging")
	errlog := flag.Bool("errlog", false, "Enable error logging")
	flag.Parse()

	switch {
	case *debug:
		log.SetLevel(log.DebugLevel)
	case *info:
		log.SetLevel(log.InfoLevel)
	case *errlog:
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}
}
