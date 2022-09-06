package handler

import (
	"fmt"
	"hand_held/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AccountTransactionPay(c echo.Context) error {
	req := new(model.AccountTransactionReq)
	if err := c.Bind(req); err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	resp, err := h.accountsRepo.AccountsTransactionInsert(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
