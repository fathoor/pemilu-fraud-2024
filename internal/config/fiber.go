package config

import (
	"github.com/fathoor/fraud24/internal/exception"
	"github.com/gofiber/fiber/v2"
)

func ProvideFiber() *fiber.Config {
	return &fiber.Config{
		CaseSensitive: true,
		StrictRouting: false,
		ErrorHandler:  exception.Handler,
	}
}
