package handler

import (
	"strconv"

	"desafio-lojas/model"
	"desafio-lojas/service"
	"desafio-lojas/utils"

	"github.com/labstack/echo/v4"
)

func CriarEstabelecimento(c echo.Context) error {
	var estab model.Estabelecimento
	if err := c.Bind(&estab); err != nil {
		return utils.JSONError(c, "Erro ao processar dados")
	}

	if err := service.CriarEstabelecimento(estab); err != nil {
		return utils.JSONError(c, err.Error())
	}

	return utils.JSONSuccess(c, "Estabelecimento criado com sucesso", nil)
}

func ListarEstabelecimentos(c echo.Context) error {
	estabelecimentos, err := service.ListarEstabelecimentos()
	if err != nil {
		return utils.JSONError(c, "Erro ao listar estabelecimentos")
	}

	return utils.JSONSuccess(c, "Lista de estabelecimentos", estabelecimentos)
}

func EditarEstabelecimento(c echo.Context) error {
	idParam := c.Param("id")
	idParsed, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.JSONError(c, "ID inválido")
	}
	id := uint(idParsed)

	var estab model.Estabelecimento
	if err := c.Bind(&estab); err != nil {
		return utils.JSONError(c, "Erro ao processar dados")
	}

	if err := service.EditarEstabelecimento(id, estab); err != nil {
		return utils.JSONError(c, err.Error())
	}

	return utils.JSONSuccess(c, "Estabelecimento atualizado", nil)
}

func RemoverEstabelecimento(c echo.Context) error {
	idParam := c.Param("id")
	idParsed, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.JSONError(c, "ID inválido")
	}

	var id uint = uint(idParsed)

	if err := service.RemoverEstabelecimento(id); err != nil {
		return utils.JSONError(c, "Não é possível remover: verifique se há lojas vinculadas")
	}

	return utils.JSONSuccess(c, "Estabelecimento removido", nil)
}

func BuscarEstabelecimentoPorID(c echo.Context) error {
	idParam := c.Param("id")
	idParsed, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.JSONError(c, "ID inválido")
	}
	id := uint(idParsed)

	estabelecimento, err := service.BuscarEstabelecimentoPorID(id)
	if err != nil {
		return utils.JSONError(c, "Estabelecimento não encontrado")
	}

	return utils.JSONSuccess(c, "Estabelecimento encontrado", estabelecimento)
}
