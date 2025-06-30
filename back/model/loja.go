package model

type Loja struct {
	ID                uint             `gorm:"primaryKey" json:"id"`
	NumeroLoja        string           `gorm:"not null" json:"numero_loja"`
	Nome              string           `gorm:"not null" json:"nome"`
	RazaoSocial       string           `gorm:"not null" json:"razao_social"`
	Endereco          string           `gorm:"not null" json:"endereco"`
	Cidade            string           `gorm:"not null" json:"cidade"`
	Estado            string           `gorm:"not null" json:"estado"`
	CEP               string           `gorm:"not null" json:"cep"`
	Numero            string           `gorm:"not null" json:"numero"`
	EstabelecimentoID uint             `gorm:"not null" json:"estabelecimento_id"`
	Estabelecimento   *Estabelecimento `gorm:"foreignKey:EstabelecimentoID" json:"estabelecimento,omitempty"`
}
