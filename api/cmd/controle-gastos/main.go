package main

import (
	"github.com/MarcoTulioFM/controle-gastos/api/config"
	"github.com/MarcoTulioFM/controle-gastos/api/router"
)

func main() {
	config.Config()
	router.Initializer()
}
