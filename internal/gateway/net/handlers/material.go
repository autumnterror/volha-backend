package handlers

import (
	"net/http"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
)

// CreateMaterial godoc
// @Summary Создать материал
// @Description Добавляет новый материал
// @Tags material
// @Accept json
// @Produce json
// @Param material body productsRPC.Material true "Новый материал"
// @Success 200 {object} views.SWGId "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 409 {object} views.SWGError "уже существует"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/material [post]
func (a *Apis) CreateMaterial(c echo.Context) error {
	const op = "handlers.CreateMaterial"
	log.Blue(op)

	if r, code, err := a.create(c, op, views.Material); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// UpdateMaterial godoc
// @Summary Обновить материал
// @Description Обновляет данные существующего материала
// @Tags material
// @Accept json
// @Produce json
// @Param material body productsRPC.Material true "Обновлённые данные материала"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/material [put]
func (a *Apis) UpdateMaterial(c echo.Context) error {
	const op = "handlers.UpdateMaterial"
	log.Blue(op)

	if r, code, err := a.update(c, op, views.Material); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// DeleteMaterial godoc
// @Summary Удалить материал
// @Description Удаляет материал по ID
// @Tags material
// @Produce json
// @Param id query string true "ID материала"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/material [delete]
func (a *Apis) DeleteMaterial(c echo.Context) error {
	const op = "handlers.DeleteMaterial"
	log.Blue(op)

	if r, code, err := a.delete(c, op, views.Material); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetAllMaterials godoc
// @Summary Получить все материалы
// @Description Возвращает список всех материалов
// @Tags material
// @Produce json
// @Success 200 {object} []productsRPC.Material "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/material/all [get]
func (a *Apis) GetAllMaterials(c echo.Context) error {
	const op = "handlers.GetAllMaterials"
	log.Blue(op)

	if r, code, err := a.getAll(c, op, views.Material); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetMaterial rand godoc
// @Summary Получить материал
// @Description Возвращает материал по id
// @Tags material
// @Produce json
// @Param id query string true "ID материал"
// @Success 200 {object} productsRPC.Material "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/material [get]
func (a *Apis) GetMaterial(c echo.Context) error {
	const op = "handlers.GetMaterial"
	log.Blue(op)

	if r, code, err := a.get(c, op, views.Material); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

