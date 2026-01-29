package main

import (
	"ProjetoNota/config"
	"ProjetoNota/router"
	"fmt"
)

func main() {

	err := config.InitializeConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}
