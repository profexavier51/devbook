package controllers

import (
	"api/src/banco"
	"api/src/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var usuario models.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()

	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte("Criando um Usuário! "))
}
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuário! "))
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuário! "))
}
func AlterarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um Usuário! "))
}
func DeletarrUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Apagando um Usuário! "))
}
