package repositories

import (
	"api/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repository usuarios) Criar(usuario models.Usuario) (uint64, error) {
	query := fmt.Sprintf("INSERT INTO usuarios (nome, nick, email, senha) VALUES ('%s', '%s', '%s', '%s')",
		usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	_, erro := repository.db.Exec(query)
	if erro != nil {
		return 0, fmt.Errorf("erro ao inserir usuário no banco de dados: %s", erro.Error())
	}

	var idInserido uint64
	err := repository.db.QueryRow("SELECT MAX(id) FROM usuarios").Scan(&idInserido)
	if err != nil {
		return 0, fmt.Errorf("erro ao obter o ID do usuário inserido: %s", err.Error())
	}

	return idInserido, nil
}
