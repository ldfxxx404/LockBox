package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var Limiter = limiter.Config{
	Expiration: 10 * time.Second,
	Max:        FiberLimitReq,
}
