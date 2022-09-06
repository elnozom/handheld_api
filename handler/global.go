package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Validate(c echo.Context) error {
	return c.JSON(http.StatusOK, "working")
}
