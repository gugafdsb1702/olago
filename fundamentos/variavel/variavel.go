package main

import "fmt"

var (
	Endereco, Telefone string
	quantidade int //quantidade = 0
	comprou bool //comprou = false
	valor float64 //valor = 0.00
	palavras rune
)

func main(){
	teste := "valor de teste"
	fmt.Printf("endereço: %s\r\n", Endereco)
	fmt.Printf("Quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v \r\n", palavras)
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
