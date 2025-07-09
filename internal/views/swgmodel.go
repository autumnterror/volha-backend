package views

type (
	// SWGProductResponse - модель продукта
	// swagger:model
	SWGProductResponse struct {
		// ID продукта
		// example: 5f8d8f9b-6f7c-4a9e-8f9b-6f7c4a9e8f9b
		ID string `json:"id"`

		// Название продукта
		// example: Смартфон Xiaomi
		Name string `json:"name"`

		// Описание продукта
		// example: Флагманский смартфон с камерой 108 МП
		Description string `json:"description"`

		// Цена продукта
		// example: 29999.99
		Price float64 `json:"price"`
	}

	// SWGProductListResponse - список продуктов
	// swagger:model
	SWGProductListResponse struct {
		// Массив продуктов
		Products []SWGProductResponse `json:"products"`
	}

	// SWGProductFilterRequest - фильтр для продуктов
	// swagger:model
	SWGProductFilterRequest struct {
		// Минимальная цена
		// example: 10000
		MinPrice float64 `json:"min_price"`

		// Максимальная цена
		// example: 50000
		MaxPrice float64 `json:"max_price"`

		// Часть названия для поиска
		// example: Xiaomi
		NamePart string `json:"name_part"`
	}

	// SWGSuccessResponse - успешный ответ
	// swagger:model
	SWGSuccessResponse struct {
		// Сообщение об успехе
		// example: Операция выполнена успешно
		Message string `json:"message"`
	}

	// SWGErrorResponse - модель ошибки
	// swagger:model
	SWGErrorResponse struct {
		// Сообщение об ошибке
		// example: Произошла ошибка
		Error string `json:"error"`
	}
)
