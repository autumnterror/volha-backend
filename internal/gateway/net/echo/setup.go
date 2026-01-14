package echo

import (
	"errors"
	"fmt"

	"github.com/autumnterror/breezynotes/pkg/utils/format"
	"github.com/autumnterror/volha-backend/internal/gateway/config"
	"github.com/autumnterror/volha-backend/internal/gateway/grpc/products"
	"github.com/autumnterror/volha-backend/internal/gateway/net/handlers"
	"github.com/autumnterror/volha-backend/internal/gateway/net/mw"
	"github.com/autumnterror/volha-backend/internal/gateway/redis"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Echo struct {
	e   *echo.Echo
	cfg *config.Config
}

const (
	ImagesDir      = "./images"
	MaxUploadBytes = 10 << 20 // 10 MB
)

func New(rds *redis.Client, a *products.Client, cfg *config.Config) *Echo {
	e := echo.New()

	h := handlers.New(a, rds, cfg)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit(fmt.Sprintf("%d", MaxUploadBytes)))
	e.Use(middleware.CORS())
	e.Static("/images", ImagesDir)

	userApi := e.Group("/api", mw.CheckId())
	{
		p := userApi.Group("/product")
		{
			p.GET("/search", h.SearchProducts)
			p.POST("/filter", h.FilterProducts)
			p.GET("/all", h.GetAllProducts)
			p.GET("", h.GetProduct)
		}

		b := userApi.Group("/brand")
		{
			b.GET("/all", h.GetAllBrands)
			b.GET("", h.GetBrand)
		}

		a := userApi.Group("/auth")
		{
			a.GET("/check", h.CheckPw)
		}

		c := userApi.Group("/category")
		{
			c.GET("/all", h.GetAllCategories)
			c.GET("", h.GetCategory)
		}

		co := userApi.Group("/color")
		{
			co.GET("/all", h.GetAllColors)
			co.GET("", h.GetColor)
		}

		m := userApi.Group("/material")
		{
			m.GET("/all", h.GetAllMaterials)
			m.GET("", h.GetMaterial)
		}

		ct := userApi.Group("/country")
		{
			ct.GET("/all", h.GetAllCountries)
			ct.GET("", h.GetCountry)
		}

		dict := userApi.Group("/dictionaries")
		{
			dict.GET("/all", h.GetAllDictionaries)
			dict.GET("/all/by-category", h.GetAllDictionariesByCategory)
		}

		cp := userApi.Group("/colorphotos")
		{
			cp.GET("/all", h.GetAllProductColorPhotos)
			cp.POST("/photos/get", h.GetPhotosByProductAndColor)
		}

		sl := userApi.Group("/slide")
		{
			sl.GET("/all", h.GetAllSlides)
			sl.GET("", h.GetSlide)
		}
		art := userApi.Group("/article")
		{
			art.GET("/all", h.GetAllArticles)
			art.GET("", h.GetArticle)
		}
	}

	adminApi := e.Group("/api", mw.CheckId(), mw.AdminAuth(cfg))
	{
		f := adminApi.Group("/files")
		{
			f.POST("", h.UploadFile)
			f.DELETE("", h.DeleteFile)
		}
		p := adminApi.Group("/product")
		{
			p.POST("", h.CreateProduct)
			p.PUT("", h.UpdateProduct)
			p.DELETE("", h.DeleteProduct)
		}

		b := adminApi.Group("/brand")
		{
			b.POST("", h.CreateBrand)
			b.PUT("", h.UpdateBrand)
			b.DELETE("", h.DeleteBrand)
		}

		c := adminApi.Group("/category")
		{
			c.POST("", h.CreateCategory)
			c.PUT("", h.UpdateCategory)
			c.DELETE("", h.DeleteCategory)
		}

		co := adminApi.Group("/color")
		{
			co.POST("", h.CreateColor)
			co.PUT("", h.UpdateColor)
			co.DELETE("", h.DeleteColor)
		}

		m := adminApi.Group("/material")
		{
			m.POST("", h.CreateMaterial)
			m.PUT("", h.UpdateMaterial)
			m.DELETE("", h.DeleteMaterial)
		}

		ct := adminApi.Group("/country")
		{
			ct.POST("", h.CreateCountry)
			ct.PUT("", h.UpdateCountry)
			ct.DELETE("", h.DeleteCountry)
		}
		cp := adminApi.Group("/colorphotos")
		{
			cp.POST("", h.CreateProductColorPhotos)
			cp.PUT("", h.UpdateProductColorPhotos)
			cp.DELETE("", h.DeleteProductColorPhotos)
		}

		sl := adminApi.Group("/slide")
		{
			sl.POST("", h.CreateSlide)
			sl.PUT("", h.UpdateSlide)
			sl.DELETE("", h.DeleteSlide)
		}
		art := adminApi.Group("/article")
		{
			art.POST("", h.CreateArticle)
			art.PUT("", h.UpdateArticle)
			art.DELETE("", h.DeleteArticle)
		}
		msg := adminApi.Group("/message")
		{
			msg.POST("", h.SendMessage)
		}
	}

	return &Echo{
		e:   e,
		cfg: cfg,
	}
}

func (e *Echo) MustRun() {
	const op = "echo.Run"

	if err := e.e.Start(fmt.Sprintf(":%d", e.cfg.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
