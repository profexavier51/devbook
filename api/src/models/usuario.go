package models

import "time"

type Usuario struct {
	ID           uint64    `json:"id,omitempty"`
	Nome         string    `json:"nome,omitempty"`
	Email        string    `json:"email,omitempty"`
	Nick		 string	   `json:"nick,omitempty"`
	Senha        string    `json:"senha,omitempty"`
	CriadoEm     time.Time `json:"criadoEm,omitempty"`
	AtualizadoEm time.Time `json:"AtualizadoEm,omitempty"`
	Ativo        bool      `json:"ativo.omitempty"`
}