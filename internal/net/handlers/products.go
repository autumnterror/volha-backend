package handlers

import (
	"context"
	"gateway/internal/grpc/products"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"time"
)

type Apis struct {
	apiProduct *products.Client
}

func New(
	apiProduct *products.Client,
) *Apis {
	return &Apis{
		apiProduct: apiProduct,
	}
}

// GetAllFilter godoc
// @Summary Получить продукты по фильтру
// @Description Возвращает список продуктов, соответствующих заданным критериям фильтрации
// @Tags products
// @Accept json
// @Produce json
// @Param filter body views.ProductFilter true "Параметры фильтрации"
// @Success 200 {object} views.SWGProductListResponse "Успешный запрос"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат запроса"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/products/getallfilter [post]
func (a *Apis) GetAllFilter(c echo.Context) error {
	const op = "handlers.GetAllFilter"

	ctx, done := context.WithTimeout(context.Background(), 3*time.Second)
	defer done()

	var filter views.ProductFilter
	if err := c.Bind(&filter); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}

	pl, err := a.apiProduct.GetAllFilter(ctx, &filter)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Internal:
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "bad on db"})
			default:
				log.Println(format.Error(op, err))
				return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
			}
		}
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
	}

	if len(pl) == 0 {
		return c.JSON(http.StatusOK, []views.Product{})
	}

	return c.JSON(http.StatusOK, pl)
}

// GetAll godoc
// @Summary Получить все продукты
// @Description Возвращает полный список всех доступных продуктов
// @Tags products
// @Produce json
// @Success 200 {object} views.SWGProductListResponse "Успешный запрос"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/products/getall [get]
func (a *Apis) GetAll(c echo.Context) error {
	const op = "handlers.GetAll"

	ctx, done := context.WithTimeout(context.Background(), 3*time.Second)
	defer done()

	pl, err := a.apiProduct.GetAll(ctx)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Internal:
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "bad on db"})
			default:
				log.Println(format.Error(op, err))
				return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
			}
		}
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
	}

	if len(pl) == 0 {
		return c.JSON(http.StatusOK, []views.Product{})
	}

	return c.JSON(http.StatusOK, pl)
}

// Create godoc
// @Summary Создать новый продукт
// @Description Добавляет новый продукт в систему
// @Tags products
// @Accept json
// @Produce json
// @Param product body views.Product true "Данные нового продукта"
// @Success 200 {object} views.SWGSuccessResponse "Продукт успешно создан"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат данных"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/products/create [post]
func (a *Apis) Create(c echo.Context) error {
	const op = "handlers.Create"

	var p views.Product
	if err := c.Bind(&p); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	p.Id = uuid.NewString()

	ctx, done := context.WithTimeout(context.Background(), 3*time.Second)
	defer done()

	err := a.apiProduct.Create(ctx, &p)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Internal:
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "bad on db"})
			default:
				log.Println(format.Error(op, err))
				return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
			}
		}
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"answer": "user create successfully",
	})
}

// Update godoc
// @Summary Обновить продукт
// @Description Обновляет информацию о существующем продукте
// @Tags products
// @Accept json
// @Produce json
// @Param id query string true "ID продукта"
// @Param product body views.Product true "Новые данные продукта"
// @Success 200 {object} views.SWGSuccessResponse "Продукт успешно обновлен"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID или формат данных"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/products/update [put]
func (a *Apis) Update(c echo.Context) error {
	const op = "handlers.Update"

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad id on query"})
	}

	var p views.Product
	if err := c.Bind(&p); err != nil {
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad JSON"})
	}
	p.Id = id

	ctx, done := context.WithTimeout(context.Background(), 3*time.Second)
	defer done()

	err := a.apiProduct.Update(ctx, &p)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Internal:
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "bad on db"})
			default:
				log.Println(format.Error(op, err))
				return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
			}
		}
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"answer": "user update successfully",
	})
}

// Delete godoc
// @Summary Удалить продукт
// @Description Удаляет продукт из системы по указанному ID
// @Tags products
// @Produce json
// @Param id query string true "ID продукта"
// @Success 200 {object} views.SWGSuccessResponse "Продукт успешно удален"
// @Failure 400 {object} views.SWGErrorResponse "Неверный ID продукта"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Failure 502 {object} views.SWGErrorResponse "Ошибка взаимодействия с сервисом"
// @Router /api/products/delete [delete]
func (a *Apis) Delete(c echo.Context) error {
	const op = "handlers.Delete"

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad id on query"})
	}

	ctx, done := context.WithTimeout(context.Background(), 3*time.Second)
	defer done()

	err := a.apiProduct.Delete(ctx, id)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Internal:
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "bad on db"})
			default:
				log.Println(format.Error(op, err))
				return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
			}
		}
		log.Println(format.Error(op, err))
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "check logs gateway"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"answer": "product delete successfully",
	})
}
