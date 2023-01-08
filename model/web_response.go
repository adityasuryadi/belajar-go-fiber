package model

import "github.com/gofiber/fiber/v2"

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(ctx *fiber.Ctx, data interface{}) {
	response := WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	}
	ctx.Status(fiber.StatusOK).JSON(response)
}

func NotFoundResponse(ctx *fiber.Ctx, data interface{}) {
	response := WebResponse{
		Code:   fiber.StatusNotFound,
		Status: "NOT_FOUND",
		Data:   data,
	}
	ctx.Status(fiber.StatusNotFound).JSON(response)
}

func InternalServerErrorResponse(ctx *fiber.Ctx, data interface{}) {
	response := WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   data,
	}
	ctx.Status(fiber.StatusInternalServerError).JSON(response)
}

func BadRequestResponse(ctx *fiber.Ctx, data interface{}) {
	response := WebResponse{
		Code:   fiber.StatusBadRequest,
		Status: "BAD_REQUEST",
		Data:   data,
	}
	ctx.Status(fiber.StatusBadRequest).JSON(response)
}
