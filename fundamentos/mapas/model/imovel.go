package model

type Imovel struct {
	Nome  string `json:"Nome"`
	Y     int    `json:"coordenada_Y"`
	X     int    `json:"coordenada_X"`
	valor int
}

func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

func (i *Imovel) GetValor() int {
	return i.valor
}
