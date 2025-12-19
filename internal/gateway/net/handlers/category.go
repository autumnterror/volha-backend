package handlers

import (
	"net/http"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
)

// CreateCategory godoc
// @Summary Создать категорию
// @Description Добавляет новую категорию
// @Tags category
// @Accept json
// @Produce json
// @Param category body productsRPC.Category true "Новая категория"
// @Success 200 {object} views.SWGId "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 409 {object} views.SWGError "уже существует"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/category [post]
func (a *Apis) CreateCategory(c echo.Context) error {
	const op = "handlers.CreateCategory"
	log.Blue(op)

	if r, code, err := a.create(c, op, views.Category); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// UpdateCategory godoc
// @Summary Обновить категорию
// @Description Обновляет информацию о категории
// @Tags category
// @Accept json
// @Produce json
// @Param category body productsRPC.Category true "Обновлённая категория"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/category [put]
func (a *Apis) UpdateCategory(c echo.Context) error {
	const op = "handlers.UpdateCategory"
	log.Blue(op)

	if r, code, err := a.update(c, op, views.Category); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// DeleteCategory godoc
// @Summary Удалить категорию
// @Description Удаляет категорию по ID
// @Tags category
// @Produce json
// @Param id query string true "ID категории"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/category [delete]
func (a *Apis) DeleteCategory(c echo.Context) error {
	const op = "handlers.DeleteCategory"
	log.Blue(op)

	if r, code, err := a.delete(c, op, views.Category); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)

		return c.JSON(code, r)
	}

}

// GetAllCategories godoc
// @Summary Получить все категории
// @Description Возвращает список всех категорий
// @Tags category
// @Produce json
// @Success 200 {object} []productsRPC.Category "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/category/all [get]
func (a *Apis) GetAllCategories(c echo.Context) error {
	const op = "handlers.GetAllCategories"
	log.Blue(op)

	if r, code, err := a.getAll(c, op, views.Category); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetCategory godoc
// @Summary Получить категорию
// @Description Возвращает категорию по id
// @Tags category
// @Produce json
// @Param id query string true "ID категории"
// @Success 200 {object} productsRPC.Category "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/category [get]
func (a *Apis) GetCategory(c echo.Context) error {
	const op = "handlers.GetCategory"
	log.Blue(op)

	if r, code, err := a.get(c, op, views.Category); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

