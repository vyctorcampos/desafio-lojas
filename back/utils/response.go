package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func JSONSuccess(c echo.Context, msg string, data interface{}) error {
	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"message": msg,
		"data":    data,
	})
}

func JSONError(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, echo.Map{
		"success": false,
		"message": msg,
	})
}
