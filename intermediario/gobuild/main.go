package main

import (
	"encoding/json"
	"fmt"

	"github.com/gugafdsb1702/olago/intermediario/gobuild/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Lucy"
	casa.X = 17
	casa.Y = 29
	casa.SetValor(100)

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))
}
