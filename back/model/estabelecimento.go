package model

type Estabelecimento struct {
	ID                    uint   `gorm:"primaryKey" json:"id"`
	NumeroEstabelecimento string `gorm:"not null" json:"numero_estabelecimento"`
	Nome                  string `gorm:"not null" json:"nome"`
	RazaoSocial           string `gorm:"not null" json:"razao_social"`
	Endereco              string `gorm:"not null" json:"endereco"`
	Cidade                string `gorm:"not null" json:"cidade"`
	Estado                string `gorm:"not null" json:"estado"`
	CEP                   string `gorm:"not null" json:"cep"`
	Numero                string `gorm:"not null" json:"numero"`
}
