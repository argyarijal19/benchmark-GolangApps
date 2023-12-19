package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ExceptionHandler(ctx *fiber.Ctx, err error) error {
	response := Response{
		Success: false,
		Code:    fiber.StatusInternalServerError,
		Message: err.Error(),
		Data:    nil,
	}

	if fail, ok := err.(*fiber.Error); ok {
		response.Code = fail.Code
		response.Message = fail.Message
	}
	return ctx.Status(response.Code).JSON(response)
}

func SendResponse(ctx *fiber.Ctx, statusCode int, success bool, message string, data interface{}) error {
	response := Response{
		Success: success,
		Code:    statusCode,
		Message: message,
		Data:    data,
	}

	return ctx.Status(statusCode).JSON(response)
}
