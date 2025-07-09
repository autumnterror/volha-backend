package main

import (
	"log"
	"os"
	"os/signal"
	"productService/config"
	"productService/internal/grpc"
	"productService/internal/pkg/psql"
	"syscall"
)

func main() {
	const op = "main"

	cfg := config.MustSetup()

	db := psql.MustConnect(cfg)
	api := psql.Driver{Driver: db.Driver}

	a := grpc.New(cfg, api)
	go a.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	a.Stop()
	if err := db.Disconnect(); err != nil {
		log.Println(err)
	}

	log.Println(op, "stop signal", sign)
}
