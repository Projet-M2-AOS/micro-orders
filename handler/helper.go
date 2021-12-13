package handler

import "github.com/gofiber/fiber/v2"

func makeErrorMsg(statusCode int, message error, errorMsg string) fiber.Map {
	return fiber.Map{
		"statusCode": statusCode,
		"message":    message,
		"error":      errorMsg,
	}
}
