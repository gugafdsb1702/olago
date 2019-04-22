package repo

import "gopkg.in/mgo.v2"

var SessaoMongo *mgo.Session

func AbreSessaoComMongo() (err error) {
	err = nil
	SessaoMongo, err = mgo.Dial("mongodb://localhost:27017/cursodego")
	return
}
