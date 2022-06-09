package handler

import (
	"hand_held/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	req := new(model.LoginReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var resp int
	err := h.db.Raw("EXEC EmployeeLogin @EmpCode = ? , @EmpPassword = ?", req.Username, req.Password).Row().Scan(&resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
