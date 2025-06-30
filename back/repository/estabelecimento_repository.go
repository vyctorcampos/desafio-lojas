package repository

import (
	"desafio-lojas/database"
	"desafio-lojas/model"
	"errors"
)

func CriarEstabelecimento(estab model.Estabelecimento) error {
	return database.DB.Create(&estab).Error
}

func ListarEstabelecimentos() ([]model.Estabelecimento, error) {
	var estabelecimentos []model.Estabelecimento
	err := database.DB.Find(&estabelecimentos).Error
	return estabelecimentos, err
}

func AtualizarEstabelecimento(id uint, estab model.Estabelecimento) error {
	estab.ID = id
	return database.DB.Model(&model.Estabelecimento{}).Where("id = ?", id).Updates(estab).Error
}

func RemoverEstabelecimento(id uint) error {
	var count int64
	err := database.DB.Model(&model.Loja{}).Where("estabelecimento_id = ?", id).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("não é possível remover: há lojas vinculadas")
	}
	return database.DB.Delete(&model.Estabelecimento{}, id).Error
}

func GetEstabelecimentoByID(id uint) (*model.Estabelecimento, error) {
	var estab model.Estabelecimento
	err := database.DB.First(&estab, id).Error
	if err != nil {
		return nil, err
	}
	return &estab, nil
}
