package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Nota struct {
	gorm.Model
	Id        int
	Data      time.Time
	Pontuacao int
	Texto     string
}
