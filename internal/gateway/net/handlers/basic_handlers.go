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

// CreateProduct godoc
// @Summary Создать продукт
// @Description Добавляет новый продукт
// @Tags product
// @Accept json
// @Produce json
// @Param material body productsRPC.ProductId true "Новый продукт"
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
// @Param material body productsRPC.ProductId true "Обновлённые данные продукта"
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

	if r, code, err := a.delete(c, op, views.Product); code != http.StatusOK {
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
// @Success 200 {object} []productsRPC.Slide "успех"
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
