package controller

import "github.com/gofiber/fiber/v2"

type FraudController interface {
	Route(app *fiber.App)
	GetFraud(ctx *fiber.Ctx) error
	UpdateFraud(ctx *fiber.Ctx) error
	GetWilayah(ctx *fiber.Ctx) error
}
