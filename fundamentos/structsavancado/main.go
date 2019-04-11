package main

import (
	"encoding/json"
	"fmt"

	"github.com/gugafdsb1702/olago/fundamentos/structsavancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Nova"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	fmt.Printf("Casa é: %v\r\n", casa)
	fmt.Print("O valor da casa é: ", casa.GetValor())

	objJson, _ := json.Marshal(casa)
	fmt.Println("A casa em Json", string(objJson))
}
