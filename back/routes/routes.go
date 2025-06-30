package routes

import (
	"desafio-lojas/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/login", handler.LoginSimulado)

	// Estabelecimentos
	e.POST("/estabelecimentos", handler.CriarEstabelecimento)
	e.GET("/estabelecimentos", handler.ListarEstabelecimentos)
	e.PUT("/estabelecimentos/:id", handler.EditarEstabelecimento)
	e.DELETE("/estabelecimentos/:id", handler.RemoverEstabelecimento)
	e.GET("/estabelecimentos/:id", handler.BuscarEstabelecimentoPorID)

	// Lojas
	e.POST("/lojas", handler.CriarLoja)
	e.GET("/lojas", handler.ListarLojas)
	e.GET("/estabelecimentos/:id/lojas", handler.ListarLojasPorEstabelecimento)
	e.PUT("/lojas/:id", handler.EditarLoja)
	e.DELETE("/lojas/:id", handler.RemoverLoja)
	e.GET("/lojas/:id", handler.BuscarLojaPorID)

}
