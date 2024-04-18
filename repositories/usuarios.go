package repositories

import (
	"api/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (u usuarios) Criar(usuarios models.Usuario) (uint64, error) {
	return 0, nil
}
