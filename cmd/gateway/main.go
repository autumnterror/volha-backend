package main

import (
	_ "github.com/autumnterror/volha-backend/docs"
	"github.com/autumnterror/volha-backend/internal/gateway/config"
	"github.com/autumnterror/volha-backend/internal/gateway/grpc/products"
	"github.com/autumnterror/volha-backend/internal/gateway/net/echo"
	"github.com/autumnterror/volha-backend/internal/gateway/redis"

	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Volha gateway REST API
// @version 0.1

// @contact.name Alex "bustard" Provor
// @contact.url https://breezy.su
// @contact.email help@breezy.su

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

	rds := redis.MustNew(cfg)

	e := echo.New(rds, p, cfg)
	go e.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	if err := e.Stop(); err != nil {
		log.Println(err)
	}

	log.Println(op, "stop signal", sign)
}
