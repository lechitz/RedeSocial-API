package controllers

import "net/http"

//CriarUsuario insere um novo usuário no Banco de Dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário."))
}

//BuscarUsuarios lista todos os usuários no Banco de Dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários."))
}

//BuscarUsuario localiza um usuário específico no Banco de Dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário."))
}

//AtualizarUsuario atualiza um usuário específico no Banco de Dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário."))
}

//DeletarUsuario exclui um usuário específico do Banco de Dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário."))
}