package handlers

import (
	"net/http"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
)

// CreateProductColorPhotos godoc
// @Summary Создать фотографии продукта для цвета
// @Description Добавляет массив фотографий для конкретного продукта и цвета
// @Tags product_color_photos
// @Accept json
// @Produce json
// @Param photos body productsRPC.ProductColorPhotos true "Новый набор фотографий"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 409 {object} views.SWGError "уже существует"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/colorphotos [post]
func (a *Apis) CreateProductColorPhotos(c echo.Context) error {
	const op = "handlers.CreateProductColorPhotos"
	log.Blue(op)

	if r, code, err := a.create(c, op, views.ProductColorPhotos); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// UpdateProductColorPhotos godoc
// @Summary Обновить фотографии продукта для цвета
// @Description Обновляет массив фотографий для конкретного продукта и цвета
// @Tags product_color_photos
// @Accept json
// @Produce json
// @Param photos body productsRPC.ProductColorPhotos true "Измененный набор фотографий"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/colorphotos [put]
func (a *Apis) UpdateProductColorPhotos(c echo.Context) error {
	const op = "handlers.UpdateProductColorPhotos"
	log.Blue(op)

	if r, code, err := a.update(c, op, views.ProductColorPhotos); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetAllProductColorPhotos godoc
// @Summary Получить все фотографии продуктов по цветам
// @Description Возвращает все записи product_color_photos
// @Tags product_color_photos
// @Produce json
// @Success 200 {object} []productsRPC.ProductColorPhotos "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/colorphotos/all [get]
func (a *Apis) GetAllProductColorPhotos(c echo.Context) error {
	const op = "handlers.GetAllProductColorPhotos"
	log.Blue(op)

	if r, code, err := a.getAll(c, op, views.ProductColorPhotos); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

