package handlers

import (
	"errors"

	"github.com/autumnterror/breezynotes/pkg/log"

	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var allowedMIMEs = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/webp": ".webp",
	"image/gif":  ".gif",
}

const (
	ImagesDir      = "./images"
	MaxUploadBytes = 10 << 20 // 10 MB
)

// UploadFile godoc
// @Summary Загрузить изображение
// @Description Сохраняет изображение в папке сервера images/. В поле "img" передается файл с расширениями jpg, png, webp, gif
// @Tags files
// @Produce json
// @Success 200 {object} views.SWGFileUploadResponse "Цвет успешно создан"
// @Failure 400 {object} views.SWGError "Неверный формат данных"
// @Failure 500 {object} views.SWGError "Ошибка на сервере"
// @Router /api/files [post]
func (a *Apis) UploadFile(c echo.Context) error {
	const op = "handlers.UploadFile"
	log.Blue(op)
	fileHeader, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "file field 'img' is required"})
	}

	src, err := fileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "cannot open uploaded file"})
	}
	defer src.Close()

	buff := make([]byte, 512)
	n, _ := io.ReadFull(src, buff)
	//to start
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return c.JSON(http.StatusInternalServerError, views.SWGError{Error: "cannot seek file"})
	}

	detected := http.DetectContentType(buff[:n])
	ext, ok := allowedMIMEs[detected]
	if !ok {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "unsupported file type"})
	}

	filename := uuid.NewString() + ext
	dstPath := filepath.Join(ImagesDir, filename)

	if err := saveMultipartFile(src, dstPath, fileHeader); err != nil {
		var msg string
		if errors.Is(err, errTooLarge) {
			msg = "file is too large"
		} else {
			msg = "failed to save file"
		}
		log.Error(op, "", err)
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: msg})
	}

	log.Green(op)
	return c.JSON(http.StatusOK, views.SWGFileUploadResponse{
		Name: filename,
	})
}

var errTooLarge = errors.New("too large")

func saveMultipartFile(src multipart.File, dstPath string, fh *multipart.FileHeader) error {
	if fh.Size > MaxUploadBytes {
		return errTooLarge
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		_ = os.Remove(dstPath)
		return err
	}

	if err := os.Chmod(dstPath, 0644); err != nil {
		return err
	}

	return nil
}

// DeleteFile godoc
// @Summary Удалить файл
// @Description Удаляет файл. Требуется передовать только имя. Пример: example.gif
// @Tags files
// @Produce json
// @Param title query string true "название файла"
// @Success 200 {object} views.SWGMessage "файл успешно удалён"
// @Failure 400 {object} views.SWGError "Неверное название"
// @Failure 500 {object} views.SWGError "Ошибка на сервере"
// @Router /api/files [delete]
func (a *Apis) DeleteFile(c echo.Context) error {
	const op = "handlers.DeleteFile"
	log.Blue(op)

	filename := c.QueryParam("title")
	if filename == "" {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "empty filename"})
	}

	if strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		return c.JSON(http.StatusBadRequest, views.SWGError{Error: "invalid filename"})
	}

	fullPath := filepath.Clean(filepath.Join(filepath.Clean(ImagesDir), filename))

	err := os.Remove(fullPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return c.JSON(http.StatusNotFound, views.SWGError{Error: "file not found"})
		}
		log.Error(op, "", err)
		return c.JSON(http.StatusInternalServerError, views.SWGError{Error: err.Error()})
	}

	log.Green(op)
	return c.JSON(http.StatusOK, views.SWGMessage{Message: "file delete successfully"})
}
