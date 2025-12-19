package handlers

import (
	"net/http"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
)

// CreateColor godoc
// @Summary Создать цвет
// @Description Добавляет новый цвет
// @Tags color
// @Accept json
// @Produce json
// @Param color body productsRPC.Color true "Новый цвет"
// @Success 200 {object} views.SWGId "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 409 {object} views.SWGError "уже существует"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/color [post]
func (a *Apis) CreateColor(c echo.Context) error {
	const op = "handlers.CreateColor"
	log.Blue(op)

	if r, code, err := a.create(c, op, views.Color); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// UpdateColor godoc
// @Summary Обновить цвет
// @Description Обновляет данные существующего цвета
// @Tags color
// @Accept json
// @Produce json
// @Param color body productsRPC.Color true "Обновлённый цвет"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/color [put]
func (a *Apis) UpdateColor(c echo.Context) error {
	const op = "handlers.UpdateColor"
	log.Blue(op)

	if r, code, err := a.update(c, op, views.Color); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// DeleteColor godoc
// @Summary Удалить цвет
// @Description Удаляет цвет по ID
// @Tags color
// @Produce json
// @Param id query string true "ID цвета"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/color [delete]
func (a *Apis) DeleteColor(c echo.Context) error {
	const op = "handlers.DeleteColor"
	log.Blue(op)

	if r, code, err := a.delete(c, op, views.Color); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetAllColors godoc
// @Summary Получить все цвета
// @Description Возвращает список всех доступных цветов
// @Tags color
// @Produce json
// @Success 200 {object} []productsRPC.Color "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/color/all [get]
func (a *Apis) GetAllColors(c echo.Context) error {
	const op = "handlers.GetAllColors"
	log.Blue(op)

	if r, code, err := a.getAll(c, op, views.Color); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetColor godoc
// @Summary Получить цвет
// @Description Возвращает цвет по id
// @Tags color
// @Produce json
// @Param id query string true "ID цвета"
// @Success 200 {object} productsRPC.Color "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/color [get]
func (a *Apis) GetColor(c echo.Context) error {
	const op = "handlers.GetBrand"
	log.Blue(op)

	if r, code, err := a.get(c, op, views.Color); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

