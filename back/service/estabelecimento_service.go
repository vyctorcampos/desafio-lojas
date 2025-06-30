package service

import (
	"desafio-lojas/model"
	"desafio-lojas/repository"
	"errors"
)

// Verifica e cria um estabelecimento
func CriarEstabelecimento(estab model.Estabelecimento) error {
	if estab.Nome == "" || estab.NumeroEstabelecimento == "" {
		return errors.New("nome e número do estabelecimento são obrigatórios")
	}
	return repository.CriarEstabelecimento(estab)
}

// Retorna todos os estabelecimentos
func ListarEstabelecimentos() ([]model.Estabelecimento, error) {
	return repository.ListarEstabelecimentos()
}

// Atualiza os dados de um estabelecimento pelo ID
func EditarEstabelecimento(id uint, estab model.Estabelecimento) error {
	return repository.AtualizarEstabelecimento(id, estab)
}

// Remove um estabelecimento se não houver lojas vinculadas
func RemoverEstabelecimento(id uint) error {
	return repository.RemoverEstabelecimento(id)
}

// Busca um estabelecimento por ID
func BuscarEstabelecimentoPorID(id uint) (*model.Estabelecimento, error) {
	return repository.GetEstabelecimentoByID(id)
}
