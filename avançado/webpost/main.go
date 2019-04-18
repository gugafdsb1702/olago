package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gugafdsb1702/olago/avançado/webpost/model"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Gustavo Barbosa"

	conteudo, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Houve um erro ao o JSON do usuário. Erro: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "https://envxcy2mhro9.x.pipedream.net", bytes.NewBuffer(conteudo))
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar um request para Google Brasil. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("name", "Yoda")
	request.Header.Set("Content-type", "Content-Type: application/json; charset=utf-8")
	reposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o serviço do request bin. Erro: ", err.Error())
		return
	}
	defer reposta.Body.Close()

	if reposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(reposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo retornado pelo request bin. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}
