package repo

import (
	"github.com/gugafdsb1702/olago/avançado/mongodb/model"
	"gopkg.in/mgo.v2/bson"
)

func ObtemLocal(codigoTelefone int) (local model.Local, err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()

	colecao := sessao.DB("cursodego").C("local")
	err = colecao.Find(bson.M{"telcode": codigoTelefone}).One(&local)
	return
}
