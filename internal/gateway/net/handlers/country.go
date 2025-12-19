package handlers

import (
	"net/http"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
)

// CreateCountry godoc
// @Summary Создать страну
// @Description Добавляет новую страну
// @Tags country
// @Accept json
// @Produce json
// @Param country body productsRPC.Country true "Новая страна"
// @Success 200 {object} views.SWGId "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 409 {object} views.SWGError "уже существует"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/country [post]
func (a *Apis) CreateCountry(c echo.Context) error {
	const op = "handlers.CreateCountry"
	log.Blue(op)

	if r, code, err := a.create(c, op, views.Country); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// UpdateCountry godoc
// @Summary Обновить страну
// @Description Обновляет данные существующей страны
// @Tags country
// @Accept json
// @Produce json
// @Param country body productsRPC.Country true "Обновлённая страна"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/country [put]
func (a *Apis) UpdateCountry(c echo.Context) error {
	const op = "handlers.UpdateCountry"
	log.Blue(op)

	if r, code, err := a.update(c, op, views.Country); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// DeleteCountry godoc
// @Summary Удалить страну
// @Description Удаляет страну по ID
// @Tags country
// @Produce json
// @Param id query string true "ID страны"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/country [delete]
func (a *Apis) DeleteCountry(c echo.Context) error {
	const op = "handlers.DeleteCountry"
	log.Blue(op)

	if r, code, err := a.delete(c, op, views.Country); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetAllCountries godoc
// @Summary Получить все страны
// @Description Возвращает список всех стран
// @Tags country
// @Produce json
// @Success 200 {object} []productsRPC.Country "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/country/all [get]
func (a *Apis) GetAllCountries(c echo.Context) error {
	const op = "handlers.GetAllCountries"
	log.Blue(op)

	if r, code, err := a.getAll(c, op, views.Country); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetCountry godoc
// @Summary Получить страну
// @Description Возвращает страну по id
// @Tags country
// @Produce json
// @Param id query string true "ID страны"
// @Success 200 {object} productsRPC.Country "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/country [get]
func (a *Apis) GetCountry(c echo.Context) error {
	const op = "handlers.GetBrand"
	log.Blue(op)

	if r, code, err := a.get(c, op, views.Country); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

