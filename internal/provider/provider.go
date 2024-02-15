package provider

import (
	"github.com/fathoor/fraud24/cmd/fraud24/controller"
	"github.com/fathoor/fraud24/cmd/fraud24/service"
	"github.com/gofiber/fiber/v2"
)

func ProvideModule(app *fiber.App) {
	ProvideFraud(app)
}

func ProvideFraud(app *fiber.App) {
	fraudService := service.ProvideFraudService()
	wilayahService := service.ProvideWilayahService()
	fraudController := controller.ProvideFraudController(&fraudService, &wilayahService)

	fraudController.Route(app)
}
