package controller

import "github.com/gofiber/fiber/v2"

type FraudController interface {
	Route(app *fiber.App)
	FraudCheck(ctx *fiber.Ctx) error
	FraudCheckCache(ctx *fiber.Ctx) error
	InitWilayah(ctx *fiber.Ctx) error
}
