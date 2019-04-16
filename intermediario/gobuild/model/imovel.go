package model

import "errors"

type Imovel struct {
	Nome  string `json:"Nome"`
	Y     int    `json:"coordenada_Y"`
	X     int    `json:"coordenada_X"`
	valor int
}

var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para um iḿovel")

var ErrValorImovelMuitoCaro = errors.New("O valor informado é muito alto.")

func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorImovelMuitoCaro
		return
	}
	i.valor = valor
	return
}

func (i *Imovel) GetValor() int {
	return i.valor
}
