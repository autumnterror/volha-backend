package main

import (
	_ "gateway/cmd/docs"
	"gateway/config"
	"gateway/copyrights"
	"gateway/internal/grpc/products"
	"gateway/internal/net/echo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Breezy notes gateway REST API
// @version 0.1

// @contact.name Alex "bustard" Provor
// @contact.url https://breezynotes.ru
// @contact.email help@breezynotes.ru

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {

	const op = "main"

	cfg := config.MustSetup()
	p, err := products.New(cfg)
	if err != nil {
		log.Panic(err)
	}

	e := echo.New(p, *cfg)
	go e.MustRun()

	if cfg.Mode != "DEV" {
		copyrights.Info()
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	if err := e.Stop(); err != nil {
		log.Println(err)
	}

	log.Println(op, "stop signal", sign)
}
