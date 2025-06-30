package repository

import (
	"desafio-lojas/database"
	"desafio-lojas/model"
)

func CriarLoja(loja model.Loja) error {
	return database.DB.Create(&loja).Error
}

func ListarLojas() ([]model.Loja, error) {
	var lojas []model.Loja
	err := database.DB.Find(&lojas).Error
	return lojas, err
}

func ListarLojasPorEstabelecimento(estabID uint) ([]model.Loja, error) {
	var lojas []model.Loja
	err := database.DB.Where("estabelecimento_id = ?", estabID).Find(&lojas).Error
	return lojas, err
}

func AtualizarLoja(id uint, loja model.Loja) error {
	loja.ID = id
	return database.DB.Model(&model.Loja{}).Where("id = ?", id).Updates(loja).Error
}

func RemoverLoja(id uint) error {
	return database.DB.Delete(&model.Loja{}, id).Error
}

func BuscarLojaComEstabelecimentoPorID(id uint) (*model.Loja, error) {
	var loja model.Loja
	err := database.DB.Preload("Estabelecimento").First(&loja, id).Error
	if err != nil {
		return nil, err
	}
	return &loja, nil
}
