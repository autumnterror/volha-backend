package main

import (
	"github.com/autumnterror/volha-backend/internal/product-service/config"
	"github.com/autumnterror/volha-backend/internal/product-service/grpc"
	"github.com/autumnterror/volha-backend/internal/product-service/psql"
	"log"
	"os"
	"os/signal"

	"syscall"
)

func main() {
	const op = "main"
	const vol = "photo color"

	cfg := config.MustSetup()

	db := psql.MustConnect(cfg)

	a := grpc.New(cfg, db.Driver)

	go a.MustRun()
	log.Printf("service started and ready to work. Vol.%s\n", vol)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	a.Stop()
	if err := db.Disconnect(); err != nil {
		log.Println(err)
	}

	log.Println(op, "stop signal", sign)
}
