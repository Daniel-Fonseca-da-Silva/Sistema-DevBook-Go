package rotas

import (
	"api/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                  "/usuarios",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarUsuario,
		RequerAutentiticacao: false,
	},
	{
		URI:                  "/usuarios",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarUsuarios,
		RequerAutentiticacao: false,
	},
	{
		URI:                  "/usuarios/{usuarioId}",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarUsuario,
		RequerAutentiticacao: false,
	},
	{
		URI:                  "/usuarios/{usuarioId}",
		Metodo:               http.MethodPut,
		Funcao:               controllers.AtualizarUsuario,
		RequerAutentiticacao: false,
	},
	{
		URI:                  "/usuarios/{usuarioId}",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarUsuario,
		RequerAutentiticacao: false,
	},
}
