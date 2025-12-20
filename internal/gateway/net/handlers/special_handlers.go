package handlers

import (
	"context"

	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"

	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// GetPhotosByProductAndColor godoc
// @Summary Получить фотографии продукта по ID продукта и цвета
// @Description Возвращает массив фотографий для указанного продукта и цвета
// @Tags product_color_photos
// @Accept json
// @Produce json
// @Param id body productsRPC.ProductColorPhotosId true "ID продукта и цвета"
// @Success 200 {object} []string "Список фотографий"
// @Failure 400 {object} views.SWGError "Неверные данные"
// @Failure 502 {object} views.SWGError "Ошибка взаимодействия с сервисом"
// @Router /api/colorphotos/photos/get [post]
func (a *Apis) GetPhotosByProductAndColor(c echo.Context) error {
	const op = "handlers.GetPhotosByProductAndColor"
	log.Blue(op)

	var pcpId productsRPC.ProductColorPhotosId
	if err := c.Bind(&pcpId); err != nil {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "bad JSON"})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	photos, err := a.apiProduct.API.GetPhotosByProductAndColor(ctx, &pcpId)
	if err != nil {
		log.Error(op, "", err)
		return c.JSON(http.StatusBadGateway, views.SWGError{Error: "failed to get photos"})
	}

	log.Green(op)
	return c.JSON(http.StatusOK, photos.GetItems())
}

// GetAllDictionariesByCategory godoc
// @Summary Получить все справочники по категории
// @Description Возвращает список всех доступных справочных данных: бренды, материалы, страны и цвета и размеры по определенной категории
// @Tags dictionaries
// @Produce json
// @Param id query string true "ID категории"
// @Success 200 {object} productsRPC.Dictionaries "Успешный запрос"
// @Failure 400 {object} views.SWGError "Ошибка на сервере"
// @Failure 502 {object} views.SWGError "Ошибка взаимодействия с сервисом"
// @Router /api/dictionaries/all/by-category [get]
func (a *Apis) GetAllDictionariesByCategory(c echo.Context) error {
	const op = "handlers.GetAllDictionariesByCategory"
	log.Blue(op)

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "missing id"})
	}

	if ds, err := a.rds.GetDictionariesByCategory(id); err == nil {
		log.Blue("read dict from cache", ds)
		log.Green(op)
		return c.JSON(http.StatusOK, ds)
	}
	log.Blue("no dict on cache")

	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	data, err := a.apiProduct.API.GetDictionaries(ctx, &productsRPC.Id{Id: id})
	if err != nil {
		log.Error(op, "", err)
		return c.JSON(http.StatusBadGateway, views.SWGError{Error: "failed to get dictionaries"})
	}

	if len(data.GetMaterials()) == 0 {
		data.Materials = []*productsRPC.Material{}
	}
	if len(data.GetColors()) == 0 {
		data.Colors = []*productsRPC.Color{}
	}
	if len(data.GetBrands()) == 0 {
		data.Brands = []*productsRPC.Brand{}
	}
	if len(data.GetCountries()) == 0 {
		data.Countries = []*productsRPC.Country{}
	}

	if err := a.rds.SetDictionariesByCategory(id, data); err != nil {
		log.Error(op, "", err)
	}
	log.Green(op)
	return c.JSON(http.StatusOK, data)
}

// GetAllDictionaries godoc
// @Summary Получить все справочники
// @Description Возвращает список всех доступных справочных данных: бренды, категории, материалы, страны и цвета
// @Tags dictionaries
// @Produce json
// @Success 200 {object} productsRPC.Dictionaries "Успешный запрос"
// @Failure 400 {object} views.SWGError "Ошибка на сервере"
// @Failure 502 {object} views.SWGError "Ошибка взаимодействия с сервисом"
// @Router /api/dictionaries/all [get]
func (a *Apis) GetAllDictionaries(c echo.Context) error {
	const op = "handlers.GetAllDictionaries"
	log.Blue(op)

	if ds, err := a.rds.GetDictionaries(); err == nil {
		log.Blue("read dict from cache", ds)
		log.Green(op)
		return c.JSON(http.StatusOK, ds)
	}
	log.Blue("no dict on cache")

	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	data, err := a.apiProduct.API.GetDictionaries(ctx, &productsRPC.Id{Id: notByCategory})
	if err != nil {
		return c.JSON(http.StatusBadGateway, views.SWGError{Error: "failed to get dictionaries"})
	}

	if len(data.GetCategories()) == 0 {
		data.Categories = []*productsRPC.Category{}
	}
	if len(data.GetMaterials()) == 0 {
		data.Materials = []*productsRPC.Material{}
	}
	if len(data.GetColors()) == 0 {
		data.Colors = []*productsRPC.Color{}
	}
	if len(data.GetBrands()) == 0 {
		data.Brands = []*productsRPC.Brand{}
	}
	if len(data.GetCountries()) == 0 {
		data.Countries = []*productsRPC.Country{}
	}

	if err := a.rds.SetDictionaries(data); err != nil {
		log.Error(op, "", err)
	}
	log.Green(op)
	return c.JSON(http.StatusOK, data)
}

// CheckPw godoc
// @Summary Проверить пароль
// @Description проверяет пароль из cookie
// @Tags auth
// @Produce json
// @Success 200 {object} views.SWGMessage "Успешный запрос"
// @Failure 401 {object} views.SWGError "Пароль не верный"
// @Router /api/auth/check [get]
func (a *Apis) CheckPw(c echo.Context) error {
	const op = "handlers.CheckPw"
	log.Blue(op)
	log.Println(op)

	cookie, err := c.Cookie("admin_pw")

	if err != nil || cookie.Value != a.cfg.AdminPW {
		return c.JSON(http.StatusUnauthorized, views.SWGError{
			Error: "unauthorized",
		})
	}
	log.Green(op)
	return c.JSON(http.StatusOK, views.SWGMessage{
		Message: "pw cool",
	})
}

// DeleteProductColorPhotos godoc
// @Summary Удалить фотографии продукта для цвета
// @Description Удаляет записи фотографий продукта для указанного цвета
// @Tags product_color_photos
// @Accept json
// @Produce json
// @Param id body productsRPC.ProductColorPhotosId true "ID продукта и цвета"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/colorphotos [delete]
func (a *Apis) DeleteProductColorPhotos(c echo.Context) error {
	const op = "handlers.DeleteProductColorPhotos"
	log.Blue(op)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	_ = a.rds.CleanDictionaries()
	var pcpId productsRPC.ProductColorPhotosId

	if err := c.Bind(&pcpId); err != nil {
		log.Error(op, "", err)
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "bad JSON"})
	}

	pcp, err := a.apiProduct.API.GetPhotosByProductAndColor(ctx, &pcpId)
	if err != nil {
		code, swg := a.mapGRPCError(op, err)
		return c.JSON(code, swg)
	}

	if _, err := a.apiProduct.API.DeleteProductColorPhotos(ctx, &pcpId); err != nil {
		code, swg := a.mapGRPCError(op, err)
		return c.JSON(code, swg)
	}

	for _, filename := range pcp.Items {
		err := deleteFile(filename)
		if err != nil {
			log.Error(op, "delete photo", err)
		}
	}

	log.Green(op)
	return c.JSON(http.StatusOK, views.SWGMessage{Message: "product color photos"})
}

// DeleteProduct godoc
// @Summary Удалить продукт
// @Description Удаляет продукт по ID
// @Tags product
// @Produce json
// @Param id query string true "ID продукта"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/product [delete]
func (a *Apis) DeleteProduct(c echo.Context) error {
	const op = "handlers.DeleteProduct"
	log.Blue(op)

	r, code, err := a.get(c, op, views.Product)
	if code != http.StatusOK {
		return c.JSON(code, err)
	}

	pr, ok := r.(*productsRPC.Product)
	if !ok {
		return c.JSON(http.StatusInternalServerError, views.SWGError{Error: "invalid type"})
	}

	if r, code, err := a.delete(c, op, views.Product); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		for _, filename := range pr.GetPhotos() {
			err := deleteFile(filename)
			if err != nil {
				log.Error(op, "delete photo", err)
			}
		}
		log.Green(op)
		return c.JSON(code, r)
	}
}
