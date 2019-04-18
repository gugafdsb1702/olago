package manipulador

import (
	"fmt"
	"net/http"

	"github.com/gugafdsb1702/olago/avançado/bdsql/model"
)

func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na renderização da página", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execuçaõ do modelo", err.Error())
	}
}
