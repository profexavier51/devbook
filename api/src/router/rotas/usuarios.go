package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasusuarios = []Rota{
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuarios,
		RequerAutenticação: false,
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticação: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticação: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AlterarUsuarios ,
		RequerAutenticação: false,
	},
	{
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarrUsuarios,
		RequerAutenticação: false,
	},
}
