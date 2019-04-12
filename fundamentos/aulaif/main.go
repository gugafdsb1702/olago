package main

import "fmt"

func main() {
	meses := 6
	situacao := true
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}
	if situacao {
		fmt.Println("Ele está devendo")
	}
	if !situacao {
		fmt.Println("Ele está adiplente")
	}
	if cidade == "Teste" {
		fmt.Println("O cliente vive na cidade Teste")
	}
	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente?", descricao)
	}
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo"
		return
	}
	descricao = "o cliente está em dia"
	return
}
