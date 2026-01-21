package handler

import (
	"fmt"
	"time"
)

type CreateNotaRequest struct {
	Data      time.Time `json:"data"`
	Pontuacao *int      `json:"pontuacao"`
	Texto     string    `json:"texto"`
}

func (c *CreateNotaRequest) Validate() error {
	if c.Data.IsZero() && c.Pontuacao == nil && c.Texto == "" {
		return fmt.Errorf("Body da request é vazio ou mal formado")
	}
	if c.Data.IsZero() {
		return errRequiredParam("Data", "time")
	}
	if c.Pontuacao == nil {
		return errRequiredParam("Pontuacao", "time")
	}
	if c.Texto == "" {
		return errRequiredParam("Texto", "time")
	}
	return nil
}

func errRequiredParam(param, typ string) error {
	return fmt.Errorf("%s (%s) is required", param, typ)
}
