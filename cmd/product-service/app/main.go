package main

import (
	"github.com/autumnterror/volha-backend/internal/product-service/api"
	"github.com/autumnterror/volha-backend/internal/product-service/config"
	"github.com/autumnterror/volha-backend/internal/product-service/infra/psql"
	"github.com/autumnterror/volha-backend/internal/product-service/repository"
	"github.com/autumnterror/volha-backend/internal/product-service/service"
	"log"
	"os"
	"os/signal"

	"syscall"
)

func main() {
	const op = "main"

	cfg := config.MustSetup()
	//DATABASE
	db := repository.MustConnect(cfg)
	rp := psql.NewRepoProvider(db.Driver)
	tx := psql.NewTxRunner(db.Driver)

	//SERVICE
	s := service.NewProductsService(tx, rp)

	//TRANSPORT
	a := api.New(cfg, s)

	go a.MustRun()
	log.Printf("service started and ready to work\n")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	a.Stop()
	if err := db.Disconnect(); err != nil {
		log.Println(err)
	}

	log.Println(op, "stop signal", sign)
}
