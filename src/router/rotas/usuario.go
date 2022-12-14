package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao:controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioId}/seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/parardeseguir",
		Metodo: http.MethodPost,
		Funcao: controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/seguidores",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/seguindo",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/atualizarsenha",
		Metodo: http.MethodPost,
		Funcao: controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}