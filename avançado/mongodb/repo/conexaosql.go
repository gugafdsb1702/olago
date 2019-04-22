package repo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1)/cursodego")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
