package provider

import (
	"github.com/fathoor/pemilu-fraud-2024/cmd/controller"
	"github.com/fathoor/pemilu-fraud-2024/cmd/service"
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
