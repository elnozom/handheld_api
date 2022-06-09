package handler

import (
	"hand_held/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DevicesCreate(c echo.Context) error {
	req := new(model.DevicesInsertReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var resp model.DeviceResponse
	err := h.db.Raw("EXEC ComUseCreate  @Imei = ? ,@ComName = ?", req.Imei, req.ComName).Row().Scan(&resp.Imei, &resp.ComName, &resp.Capital)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeviceFind(c echo.Context) error {
	req := new(model.DevicesFindReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var resp model.DeviceResponse
	err := h.db.Raw("EXEC ComUseFind  @Imei = ?", req.Imei).Row().Scan(&resp.Imei, &resp.ComName, &resp.Capital)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
