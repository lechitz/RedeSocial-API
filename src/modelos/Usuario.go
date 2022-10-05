package modelos

import (
	"errors"
	"strings"
	"time"
)

//Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

//Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatório")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome) //Removendo os espaços nas extremidades do nome
	usuario.Nick = strings.TrimSpace(usuario.Nick) //Removendo os espaços nas extremidades do nick
	usuario.Email = strings.TrimSpace(usuario.Email) //Removendo os espaços nas extremidades do email
}