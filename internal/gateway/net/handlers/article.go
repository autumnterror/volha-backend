package handlers

import (
	"net/http"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
)

// CreateArticle godoc
// @Summary Создать новую статью
// @Description Добавляет новую статью. Id и CreationTime создаются автоматически. Можно передавать значения по умолчанию
// @Tags blog
// @Accept json
// @Produce json
// @Param article body productsRPC.Article true "Данные новой статьи"
// @Success 200 {object} views.SWGId "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 409 {object} views.SWGError "уже существует"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/article [post]
func (a *Apis) CreateArticle(c echo.Context) error {
	const op = "handlers.CreateArticle"
	log.Blue(op)

	if r, code, err := a.create(c, op, views.Article); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// UpdateArticle godoc
// @Summary Обновить статью
// @Description Обновляет информацию о статье по ID
// @Tags blog
// @Accept json
// @Produce json
// @Param article body productsRPC.Article true "Данные статьи"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/article [put]
func (a *Apis) UpdateArticle(c echo.Context) error {
	const op = "handlers.UpdateArticle"
	log.Blue(op)

	if r, code, err := a.update(c, op, views.Article); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// DeleteArticle godoc
// @Summary Удалить статью
// @Description Удаляет статью по ID
// @Tags blog
// @Produce json
// @Param id query string true "ID статью"
// @Success 200 {object} views.SWGMessage "успех"
// @Failure 400 {object} views.SWGError "неправильный ввод"
// @Failure 404 {object} views.SWGError "не найдено"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/article [delete]
func (a *Apis) DeleteArticle(c echo.Context) error {
	const op = "handlers.DeleteArticle"
	log.Blue(op)

	if r, code, err := a.delete(c, op, views.Article); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetAllArticles godoc
// @Summary Получить все статьи
// @Description Возвращает список статей
// @Tags blog
// @Produce json
// @Success 200 {object} []productsRPC.Article "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/article/all [get]
func (a *Apis) GetAllArticles(c echo.Context) error {
	const op = "handlers.GetAllArticles"
	log.Blue(op)

	if r, code, err := a.getAll(c, op, views.Article); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}

// GetArticle godoc
// @Summary Получить статью
// @Description Возвращает статью по id
// @Tags blog
// @Produce json
// @Param id query string true "ID статью"
// @Success 200 {object} productsRPC.Article "успех"
// @Failure 502 {object} views.SWGError "ошибка в микросервисе"
// @Router /api/article [get]
func (a *Apis) GetArticle(c echo.Context) error {
	const op = "handlers.GetArticle"
	log.Blue(op)

	if r, code, err := a.get(c, op, views.Article); code != http.StatusOK {
		return c.JSON(code, err)
	} else {
		log.Green(op)
		return c.JSON(code, r)
	}
}
