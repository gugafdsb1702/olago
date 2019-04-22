package manipulador

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gugafdsb1702/olago/avançado/mongodb/model"
	"github.com/gugafdsb1702/olago/avançado/mongodb/repo"
)

func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um número válido. Verifique.", http.StatusBadRequest)
		fmt.Println("[Local] erro ao converter o número enviado: ", err.Error())
		return
	}
	sql := "select country, city, telcode from place where telcode = ?"
	linha, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível pesquisar esse número", http.StatusInternalServerError)
		fmt.Println("[Local] não foi possível executar a query: ", sql, "Erro: ", err.Error())
		return
	}
	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possivel pesquisar esse número.", http.StatusInternalServerError)
			fmt.Println("[Local] Não foi fazer o binding dos dados do banco na Struct, erro: ", err.Error())
			return
		}
	}
	localMongo, err := repo.ObtemLocal(codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível pesquisar esse número.", http.StatusInternalServerError)
		fmt.Println("[local] não foi possível executar a query no MongoDB. Erro: ", err.Error())
		return
	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", localMongo); err != nil {
		http.Error(w, "Houve um erro na renderização da página", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execuçaõ do modelo", err.Error())
	}
}
