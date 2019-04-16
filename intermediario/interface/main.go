package main

import (
	"fmt"

	"github.com/gugafdsb1702/olago/intermediario/interface/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	QueroAcordarComUmCacarejo(jojo)
	QueroOuvirUmaPataNoLago(jojo)
}

func QueroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func QueroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
