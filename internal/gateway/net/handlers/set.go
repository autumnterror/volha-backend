package handlers

import (
	"net/http"
	"strconv"

	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

const notByCategory = "CLR"

func classifyQuery(q string) *productsRPC.ProductSearchWithPagination {
	if _, err := xid.FromBytes([]byte(q)); err == nil {
		return &productsRPC.ProductSearchWithPagination{Id: q}
	}

	if len(q) == 8 && isDigitsOnly(q) {
		return &productsRPC.ProductSearchWithPagination{Article: q}
	}

	return &productsRPC.ProductSearchWithPagination{Title: q}
}

func isDigitsOnly(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

const successPag = 0

func pag(c echo.Context) (int, int, int, any) {
	s := c.QueryParam("start")
	if s == "" {
		return 0, 0, http.StatusBadRequest, views.SWGError{Error: "bad start"}
	}
	en := c.QueryParam("end")
	if en == "" {
		return 0, 0, http.StatusBadRequest, views.SWGError{Error: "bad end"}
	}

	start, err := strconv.Atoi(s)
	if err != nil {
		return 0, 0, http.StatusBadRequest, views.SWGError{Error: "start must be int"}
	}
	end, err := strconv.Atoi(en)
	if err != nil {
		return 0, 0, http.StatusBadRequest, views.SWGError{Error: "end must be int"}
	}

	if start > end {
		return 0, 0, http.StatusOK, []productsRPC.Product{}
	}

	if start < 0 {
		return 0, 0, http.StatusBadRequest, views.SWGError{Error: "start < 0!"}
	}
	return start, end, successPag, nil
}
