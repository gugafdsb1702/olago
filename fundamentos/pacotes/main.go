package main

import (
	"fmt"
	"github.com/gugafdsb1702/olago/pacotes/prefixo"
)
var NomeDoUsuario string = "Guuh"

func main(){
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
}
