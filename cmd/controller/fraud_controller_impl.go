package controller

import (
	"github.com/fathoor/pemilu-fraud-2024/cmd/entity"
	"github.com/fathoor/pemilu-fraud-2024/cmd/service"
	"github.com/gofiber/fiber/v2"
)

type fraudControllerImpl struct {
	fraudService   service.FraudService
	wilayahService service.WilayahService
}

func (f *fraudControllerImpl) Route(app *fiber.App) {
	app.Get("/", f.GetFraud)
	app.Get("/fraud", f.UpdateFraud)
	app.Get("/wilayah", f.GetWilayah)
}

func (f *fraudControllerImpl) GetFraud(ctx *fiber.Ctx) error {
	kota := ctx.Query("kota")

	fraud := f.fraudService.FraudCache(kota)

	return ctx.Status(fiber.StatusOK).JSON(fraud)
}

func (f *fraudControllerImpl) UpdateFraud(ctx *fiber.Ctx) error {
	kota := ctx.Query("kota")

	fraud := f.fraudService.FraudCheck(kota)

	return ctx.Status(fiber.StatusOK).JSON(fraud)
}

func (f *fraudControllerImpl) GetWilayah(ctx *fiber.Ctx) error {
	f.wilayahService.GetWilayah()

	return ctx.Status(fiber.StatusOK).JSON(entity.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "Wilayah berhasil diinisialisasi",
	})
}

func ProvideFraudController(fraudService *service.FraudService, wilayahService *service.WilayahService) FraudController {
	return &fraudControllerImpl{*fraudService, *wilayahService}
}
