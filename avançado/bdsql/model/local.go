package model

import "database/sql"

type Local struct {
	Pais             string         `json: "Pais" db: "country"`
	Cidade           sql.NullString `json: "cidade" db: "city"`
	CodigoTelefonico int            `josn: "cod_telefone" db: "telcode"`
}
