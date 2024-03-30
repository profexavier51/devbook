package repositories

import "database/sql"

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios{
	return &usuarios{db}
}