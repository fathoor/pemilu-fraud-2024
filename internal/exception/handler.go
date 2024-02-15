package exception

import (
	"errors"
	"github.com/fathoor/fraud24/cmd/fraud24/entity"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, err error) error {
	var badRequestError BadRequestError
	var notFoundError NotFoundError

	switch {
	case errors.As(err, &badRequestError):
		return c.Status(fiber.StatusBadRequest).JSON(entity.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	case errors.As(err, &notFoundError):
		return c.Status(fiber.StatusNotFound).JSON(entity.Response{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(entity.Response{
		Code:   fiber.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err.Error(),
	})
}
