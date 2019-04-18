package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a página do Google Brasil. Erro:", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do Google Brasil. Erro:", err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	request, err := http.NewRequest("GET", "https://www.nerduniverse.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request. Erro:", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(request)
	resposta, err = cliente.Get("https://www.nerduniverse.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a página do Google Brasil. Erro:", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do Google Brasil. Erro:", err.Error())
			return
		}
		fmt.Println("")
		fmt.Println(string(corpo))
	}
}
