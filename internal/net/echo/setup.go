package echo

import (
	"errors"
	"fmt"
	"gateway/config"
	"gateway/internal/grpc/products"
	"gateway/internal/net/handlers"
	"gateway/internal/utils/format"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

type Echo struct {
	e   *echo.Echo
	cfg config.Config
}

func New(a *products.Client, cfg config.Config) *Echo {
	e := echo.New()

	h := handlers.New(a)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger(), middleware.Recover())

	p := e.Group("/api/products")
	{
		p.GET("/getall", h.GetAll)
		p.POST("/getallfilter", h.GetAllFilter)

		p.POST("/create", h.Create)
		p.PUT("/update", h.Update)
		p.DELETE("/delete", h.Delete)
	}
	return &Echo{
		e:   e,
		cfg: cfg,
	}
}

func (e *Echo) MustRun() {
	const op = "echo.Run"

	if err := e.e.Start(fmt.Sprintf(":%d", e.cfg.Port)); err != nil && !errors.Is(http.ErrServerClosed, err) {
		e.e.Logger.Fatal(format.Error(op, err))
	}
}

func (e *Echo) Stop() error {
	const op = "echo.Stop"

	if err := e.e.Close(); err != nil {
		return format.Error(op, err)
	}
	return nil
}
