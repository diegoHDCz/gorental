package main

import (
	"fmt"

	"github.com/diegoHDCz/gorental/config"
	"github.com/diegoHDCz/gorental/router"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println("Deu ruim no db ", err)
		panic(err)
		return
	}
	router.Initialize()
}
