package handler

import (
	"strconv"

	"desafio-lojas/model"
	"desafio-lojas/service"
	"desafio-lojas/utils"

	"github.com/labstack/echo/v4"
)

func CriarLoja(c echo.Context) error {
	var loja model.Loja
	if err := c.Bind(&loja); err != nil {
		return utils.JSONError(c, "Erro ao processar dados de criação de loja")
	}

	if err := service.CriarLoja(loja); err != nil {
		return utils.JSONError(c, err.Error())
	}

	return utils.JSONSuccess(c, "Loja criada com sucesso", nil)
}

func ListarLojas(c echo.Context) error {
	lojas, err := service.ListarLojas()
	if err != nil {
		return utils.JSONError(c, "Erro ao listar lojas")
	}

	return utils.JSONSuccess(c, "Lista de lojas", lojas)
}

func ListarLojasPorEstabelecimento(c echo.Context) error {
	idParam := c.Param("id")
	idParsed, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.JSONError(c, "ID inválido")
	}
	id := uint(idParsed)

	lojas, err := service.ListarLojasPorEstabelecimento(id)
	if err != nil {
		return utils.JSONError(c, "Erro ao buscar lojas")
	}

	return utils.JSONSuccess(c, "Lojas do estabelecimento", lojas)
}

func EditarLoja(c echo.Context) error {
	idParam := c.Param("id")
	idParsed, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.JSONError(c, "ID inválido ao editar loja")
	}
	id := uint(idParsed)

	var loja model.Loja
	if err := c.Bind(&loja); err != nil {
		return utils.JSONError(c, "Erro ao processar dados ao editar loja")
	}

	if err := service.EditarLoja(id, loja); err != nil {
		return utils.JSONError(c, err.Error())
	}

	return utils.JSONSuccess(c, "Loja atualizada com sucesso", nil)
}

func RemoverLoja(c echo.Context) error {
	idParam := c.Param("id")
	idParsed, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.JSONError(c, "ID inválido ao remover loja")
	}
	id := uint(idParsed)

	if err := service.RemoverLoja(id); err != nil {
		return utils.JSONError(c, "Erro ao remover loja")
	}

	return utils.JSONSuccess(c, "Loja removida com sucesso", nil)
}

func BuscarLojaPorID(c echo.Context) error {
	idParam := c.Param("id")
	idParsed, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.JSONError(c, "ID inválido ao buscar unica loja")
	}
	id := uint(idParsed)

	loja, err := service.BuscarLojaPorID(id)
	if err != nil {
		return utils.JSONError(c, "Loja não encontrada")
	}

	return utils.JSONSuccess(c, "Loja encontrada", loja)
}
