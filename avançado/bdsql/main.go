package main

import (
	"fmt"
	"net/http"

	"github.com/gugafdsb1702/olago/avançado/bdsql/manipulador"
	"github.com/gugafdsb1702/olago/avançado/bdsql/repo"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor...")
}

func main() {
	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados:", err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Estou escutando comandante...")
	http.ListenAndServe(":8181", nil)
}
