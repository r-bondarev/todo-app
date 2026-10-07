package main

import (
	"log"

	"github.com/r-bondarev/todo-app/pkg/handler"
)

func main() {
    handlers := new(handler.Handler)

    srv := new(Server)
    if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
        log.Fatalf("error occurred while running http server: %s", err.Error())
    }
}
