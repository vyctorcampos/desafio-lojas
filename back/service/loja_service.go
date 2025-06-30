package service

import (
	"desafio-lojas/model"
	"desafio-lojas/repository"
	"errors"
)

// Cria uma loja
func CriarLoja(loja model.Loja) error {
	if loja.Nome == "" || loja.NumeroLoja == "" {
		return errors.New("nome e número da loja são obrigatórios")
	}
	return repository.CriarLoja(loja)
}

// Lista todas as lojas
func ListarLojas() ([]model.Loja, error) {
	return repository.ListarLojas()
}

// Lista lojas de um estabelecimento
func ListarLojasPorEstabelecimento(estabID uint) ([]model.Loja, error) {
	return repository.ListarLojasPorEstabelecimento(estabID)
}

// Edita os dados de uma loja
func EditarLoja(id uint, loja model.Loja) error {
	return repository.AtualizarLoja(id, loja)
}

// Remove uma loja
func RemoverLoja(id uint) error {
	return repository.RemoverLoja(id)
}

// Busca um loja
func BuscarLojaPorID(id uint) (*model.Loja, error) {
	return repository.BuscarLojaComEstabelecimentoPorID(id)
}
