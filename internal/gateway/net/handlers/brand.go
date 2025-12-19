package handlers

import (
	"net/http"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
)

// CreateBrand godoc
// @Summary Создать новый бренд
// @Description Добавляет новый бренд
// @Tags brands
// @Accept json
// @Produce json
// @Param brand body productsRPC.Brand true "Данные нового бренда"
// @Success 200 {object} views.SWGId "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 409 {object} views.SWGError "уже существует"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/brand [post]
func (a *Apis) CreateBrand(c echo.Context) error {
	const op = "handlers.CreateBrand"
	log.Blue(op)

	if r, code, err := a.create(c, op, views.Brand); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// UpdateBrand godoc
// @Summary Обновить бренд
// @Description Обновляет информацию о бренде по ID
// @Tags brands
// @Accept json
// @Produce json
// @Param brand body productsRPC.Brand true "Данные бренда"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/brand [put]
func (a *Apis) UpdateBrand(c echo.Context) error {
	const op = "handlers.UpdateBrand"
	log.Blue(op)

	if r, code, err := a.update(c, op, views.Brand); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// DeleteBrand godoc
// @Summary Удалить бренд
// @Description Удаляет бренд по ID
// @Tags brands
// @Produce json
// @Param id query string true "ID бренда"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/brand [delete]
func (a *Apis) DeleteBrand(c echo.Context) error {
	const op = "handlers.DeleteBrand"
	log.Blue(op)

	if r, code, err := a.delete(c, op, views.Brand); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetAllBrands godoc
// @Summary Получить все бренды
// @Description Возвращает список брендов
// @Tags brands
// @Produce json
// @Success 200 {object} []productsRPC.Brand "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/brand/all [get]
func (a *Apis) GetAllBrands(c echo.Context) error {
	const op = "handlers.GetAllBrands"
	log.Blue(op)

	if r, code, err := a.getAll(c, op, views.Brand); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetBrand godoc
// @Summary Получить бренд
// @Description Возвращает бренд по id
// @Tags brands
// @Produce json
// @Param id query string true "ID бренда"
// @Success 200 {object} productsRPC.Brand "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/brand [get]
func (a *Apis) GetBrand(c echo.Context) error {
	const op = "handlers.GetBrand"
	log.Blue(op)

	if r, code, err := a.get(c, op, views.Brand); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

