package main

import (
	"github.com/mao360/Panacea"
	"github.com/mao360/Panacea/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(Panacea.Server)
	if err := srv.Run("8000", handlers.InitRouts()); err != nil {
		log.Fatalf("error occured while running httpServer: %s", err.Error())
	}
}
