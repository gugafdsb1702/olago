package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gugafdsb1702/olago/avançado/webgo/model"
)

func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := ModeloOla.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na renderização da página", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execuçaõ do modelo", err.Error())
	}
}
