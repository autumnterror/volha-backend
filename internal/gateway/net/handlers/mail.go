package handlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/mailersend/mailersend-go"
)

// SendMessage godoc
// @Summary Отправить сообщение
// @Description
// @Tags message
// @Produce json
// @Param data body views.Email true "Новый цвет"
// @Success 200 {object} views.SWGFileUploadResponse "Цвет успешно создан"
// @Failure 400 {object} views.SWGError "Неверный формат данных"
// @Failure 500 {object} views.SWGError "Ошибка на сервере"
// @Router /api/message [post]
func (a *Apis) SendMessage(c echo.Context) error {
	const op = "handlers.UploadFile"
	log.Blue(op)

	var e views.Email
	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "bad JSON"})
	}

	if err := godotenv.Load(); err != nil {
		log.Error(op, "cant load .env", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	ms := mailersend.NewMailersend(os.Getenv("MESSAGE_API_KEY"))

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	subject := e.Subject
	text := e.Text
	html := e.Html

	from := mailersend.From{
		Name:  "Volha company",
		Email: a.cfg.Email,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  e.RecipientName,
			Email: e.RecipientEmail,
		},
	}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)

	_, err := ms.Email.Send(ctx, message)
	if err != nil {
		log.Error(op, "send", err)
		return c.JSON(http.StatusBadGateway, views.SWGError{Error: "bad API"})
	}

	log.Green(op)
	return c.JSON(http.StatusOK, nil)
}
