package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Nota struct {
	gorm.Model
	Data      time.Time
	Pontuacao int
	Texto     string
}

type NotaResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Pontuacao int       `json:"pontuacao"`
	Texto     string    `json:"texto"`
}
