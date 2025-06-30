package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginSimulado(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login realizado com sucesso!",
	})
}
