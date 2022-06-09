package model

type LoginReq struct {
	Username int    `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
