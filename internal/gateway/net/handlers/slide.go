package handlers

import (
	"net/http"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
)

// CreateSlide godoc
// @Summary Создать слайд
// @Description Добавляет новый слайд
// @Tags slide
// @Accept json
// @Produce json
// @Param material body productsRPC.Slide true "Новый слайд"
// @Success 200 {object} views.SWGId "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 409 {object} views.SWGError "уже существует"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/slide [post]
func (a *Apis) CreateSlide(c echo.Context) error {
	const op = "handlers.CreateSlide"
	log.Blue(op)

	if r, code, err := a.create(c, op, views.Slide); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// UpdateSlide godoc
// @Summary Обновить слайд
// @Description Обновляет данные существующего слайда
// @Tags slide
// @Accept json
// @Produce json
// @Param material body productsRPC.Slide true "Обновлённые данные слайда"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/slide [put]
func (a *Apis) UpdateSlide(c echo.Context) error {
	const op = "handlers.UpdateSlide"
	log.Blue(op)

	if r, code, err := a.update(c, op, views.Slide); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// DeleteSlide godoc
// @Summary Удалить слайд
// @Description Удаляет слайд по ID
// @Tags slide
// @Produce json
// @Param id query string true "ID слайда"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/slide [delete]
func (a *Apis) DeleteSlide(c echo.Context) error {
	const op = "handlers.DeleteSlide"
	log.Blue(op)

	if r, code, err := a.delete(c, op, views.Slide); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetAllSlides godoc
// @Summary Получить все слайды
// @Description Возвращает список всех слайдов
// @Tags slide
// @Produce json
// @Success 200 {object} []productsRPC.Slide "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/slide/all [get]
func (a *Apis) GetAllSlides(c echo.Context) error {
	const op = "handlers.GetAllSlides"
	log.Blue(op)

	if r, code, err := a.getAll(c, op, views.Slide); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetSlide rand godoc
// @Summary Получить слайд
// @Description Возвращает слайд по id
// @Tags slide
// @Produce json
// @Param id query string true "ID слайд"
// @Success 200 {object} productsRPC.Slide "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/slide [get]
func (a *Apis) GetSlide(c echo.Context) error {
	const op = "handlers.GetSlide"
	log.Blue(op)

	if r, code, err := a.get(c, op, views.Slide); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

