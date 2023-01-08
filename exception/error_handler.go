package exception

import (
	"go-blog/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(model.WebResponse{
				Code:   fiber.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
			})
	}

	errData := logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": string(ctx.Request().Header.Method()),
		"uri":    ctx.Request().URI().String(),
		"ip":     ctx.IP(),
		"error":  err.Error(),
	}

	// LogError(errData)
	return ctx.Status(fiber.StatusInternalServerError).
		JSON(model.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   errData,
		})
}
