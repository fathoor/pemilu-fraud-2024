package controller

import (
	"github.com/fathoor/fraud24/cmd/fraud24/entity"
	"github.com/fathoor/fraud24/cmd/fraud24/service"
	"github.com/gofiber/fiber/v2"
)

type fraudControllerImpl struct {
	fraudService   service.FraudService
	wilayahService service.WilayahService
}

func (f *fraudControllerImpl) Route(app *fiber.App) {
	app.Get("/", f.FraudCheck)
	app.Get("/cache", f.FraudCheckCache)
	app.Get("/init", f.InitWilayah)
}

func (f *fraudControllerImpl) FraudCheck(ctx *fiber.Ctx) error {
	fraud := f.fraudService.FraudCheck()

	return ctx.Status(fiber.StatusOK).JSON(fraud)
}

func (f *fraudControllerImpl) FraudCheckCache(ctx *fiber.Ctx) error {
	fraud := f.fraudService.FraudCheckCache()

	return ctx.Status(fiber.StatusOK).JSON(fraud)
}

func (f *fraudControllerImpl) InitWilayah(ctx *fiber.Ctx) error {
	f.wilayahService.InitWilayah()

	return ctx.Status(fiber.StatusOK).JSON(entity.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "Wilayah berhasil diinisialisasi",
	})
}

func ProvideFraudController(fraudService *service.FraudService, wilayahService *service.WilayahService) FraudController {
	return &fraudControllerImpl{*fraudService, *wilayahService}
}
