package main

import (
	"fmt"

	"github.com/gugafdsb1702/olago/fundamentos/structsavancado/model"
)

var cache map[string]model.Imovel

func main() {
	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa Amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(70000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Cache |%s| = %+v\n\r", chave, imovel)
	}

	_, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, "Casa Amarela")
	}

	fmt.Println("Quantos itens há no cache?", len(cache))
}
