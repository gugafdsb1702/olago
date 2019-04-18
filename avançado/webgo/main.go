package main

import (
	"fmt"
	"net/http"

	"github.com/gugafdsb1702/olago/avançado/webgo/manipulador"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Olá Mundo!\n")
		fmt.Fprint(w, "Olá Mundo!\n")
		fmt.Fprint(w, "Olá Mundo!\n")
		fmt.Fprint(w, "Olá Mundo!\n")
		fmt.Fprint(w, "Olá Mundo!\n")

	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Estou escutando, comandante")
	http.ListenAndServe(":8181", nil)
}
