package model

import (
	"time"
)

type Nota struct {
	Id        int
	Data      time.Time
	Pontuacao int
	Texto     string
}
