package handlers

import (
	"net/http"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
)

// CreateProduct godoc
// @Summary Создать продукт
// @Description Добавляет новый продукт
// @Tags product
// @Accept json
// @Produce json
// @Param product body productsRPC.ProductId true "Новый продукт"
// @Success 200 {object} views.SWGId "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 409 {object} views.SWGError "уже существует"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/product [post]
func (a *Apis) CreateProduct(c echo.Context) error {
	const op = "handlers.CreateProduct"
	log.Blue(op)

	if r, code, err := a.create(c, op, views.Product); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// UpdateProduct godoc
// @Summary Обновить продукт
// @Description Обновляет данные существующего продукта
// @Tags product
// @Accept json
// @Produce json
// @Param product body productsRPC.ProductId true "Обновлённые данные продукта"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/product [put]
func (a *Apis) UpdateProduct(c echo.Context) error {
	const op = "handlers.UpdateProduct"
	log.Blue(op)

	if r, code, err := a.update(c, op, views.Product); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetProduct rand godoc
// @Summary Получить продукт
// @Description Возвращает продукт по id
// @Tags product
// @Produce json
// @Param id query string true "ID продукт"
// @Success 200 {object} productsRPC.Product "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/product [get]
func (a *Apis) GetProduct(c echo.Context) error {
	const op = "handlers.GetProduct"
	log.Blue(op)

	if r, code, err := a.get(c, op, views.Product); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

