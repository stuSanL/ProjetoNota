package handler

import (
	"fmt"
)

func errRequiredParam(param, typ string) error {
	return fmt.Errorf("%s (%s) is required", param, typ)
}

type AddNotaRequest struct {
	Pontuacao *int   `json:"pontuacao"`
	Texto     string `json:"texto"`
}

func (c *AddNotaRequest) Validate() error {
	if c.Pontuacao == nil && c.Texto == "" {
		return fmt.Errorf("Body da request é vazio ou mal formado")
	}
	if c.Pontuacao == nil {
		return errRequiredParam("Pontuacao", "time")
	}
	if c.Texto == "" {
		return errRequiredParam("Texto", "time")
	}
	return nil
}

type UpdateNotaRequest struct {
	Pontuacao *int   `json:"pontuacao"`
	Texto     string `json:"texto"`
}

func (c *UpdateNotaRequest) Validate() error {
	if c.Pontuacao == nil && c.Texto == "" {
		return fmt.Errorf("Nenhum campo provido")
	}
	return nil
}
