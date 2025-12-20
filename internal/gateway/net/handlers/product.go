package handlers

import (
	"context"
	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"net/http"
	"time"

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

// SearchProducts godoc
// @Summary Поиск продуктов
// @Description Ищет продукт по ID (UUID), article (8 цифр) или названию (частичное совпадение)
// @Tags product
// @Produce json
// @Param prompt query string true "Поисковый запрос"
// @Param start query int true  "start > 0"
// @Param end query int true  "end"
// @Success 200 {object} []productsRPC.Product
// @Failure 400 {object} views.SWGError
// @Failure 502 {object} views.SWGError
// @Router /api/product/search [get]
func (a *Apis) SearchProducts(c echo.Context) error {
	const op = "handlers.SearchProducts"
	log.Blue(op)

	query := c.QueryParam("prompt")
	if query == "" {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "empty query"})
	}

	filter := classifyQuery(query)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	start, end, code, res := pag(c)
	if code != successPag {
		if code == http.StatusOK {
			log.Green(op)
		}
		return c.JSON(code, res)
	}

	list, err := a.apiProduct.API.SearchProducts(ctx, &productsRPC.ProductSearchWithPagination{
		Id:      filter.GetId(),
		Article: filter.GetArticle(),
		Title:   filter.GetTitle(),
		Pag: &productsRPC.Pagination{
			Start:  int32(start),
			Finish: int32(end),
		},
	})
	if err != nil {
		log.Error(op, "", err)
		return c.JSON(http.StatusBadGateway, views.SWGError{Error: "could not search products"})
	}

	items := list.GetItems()
	if len(items) == 0 {
		items = []*productsRPC.Product{}
	}

	log.Green(op)
	return c.JSON(http.StatusOK, list)
}

// FilterProducts godoc
// @Summary Получить продукты по фильтру
// @Description Возвращает список продуктов, соответствующих заданным критериям фильтрации
// @Tags product
// @Accept json
// @Produce json
// @Param filter body productsRPC.ProductFilter true "Параметры фильтрации"
// @Success 200 {object} []productsRPC.Product "Успешный запрос"
// @Failure 400 {object} views.SWGError "Неверный формат запроса"
// @Failure 500 {object} views.SWGError "Ошибка на сервере"
// @Failure 502 {object} views.SWGError "Ошибка взаимодействия с сервисом"
// @Router /api/product/filter [post]
func (a *Apis) FilterProducts(c echo.Context) error {
	const op = "handlers.FilterProducts"
	log.Blue(op)

	var f productsRPC.ProductFilter
	if err := c.Bind(&f); err != nil {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "bad JSON"})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	list, err := a.apiProduct.API.FilterProducts(ctx, &f)
	if err != nil {
		log.Error(op, "", err)
		return c.JSON(http.StatusBadGateway, views.SWGError{Error: "could not filter products"})
	}

	items := list.GetItems()
	if len(items) == 0 {
		log.Green(op)
		return c.JSON(http.StatusOK, []productsRPC.Product{})
	}

	log.Green(op)
	return c.JSON(http.StatusOK, list)
}

// GetAllProducts godoc
// @Summary Получить все продукты
// @Description Возвращает список всех продуктов
// @Tags product
// @Produce json
// @Param start query int true  "start > 0"
// @Param end query int true  "end"
// @Success 200 {object} []productsRPC.Product
// @Failure 400 {object} views.SWGError
// @Failure 502 {object} views.SWGError "Ошибка взаимодействия с сервисом"
// @Router /api/product/all [get]
func (a *Apis) GetAllProducts(c echo.Context) error {
	const op = "handlers.GetAllProducts"
	log.Blue(op)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	start, end, code, res := pag(c)
	if code != successPag {
		if code == http.StatusOK {
			log.Green(op)
		}
		return c.JSON(code, res)
	}

	list, err := a.apiProduct.API.GetAllProducts(ctx, &productsRPC.Pagination{
		Start:  int32(start),
		Finish: int32(end),
	})
	if err != nil {
		log.Error(op, "", err)
		return c.JSON(http.StatusBadGateway, views.SWGError{Error: "could not fetch products"})
	}

	items := list.GetItems()
	if len(items) == 0 {
		log.Green(op)
		return c.JSON(http.StatusOK, []productsRPC.Product{})
	}

	log.Green(op)
	return c.JSON(http.StatusOK, list)
}
