package model

type DevicesInsertReq struct {
	ComName string `json:"com_name" validate:"required"`
	Imei    string `json:"imei" validate:"required"`
}

type DevicesFindReq struct {
	Imei string `json:"imei" validate:"required"`
}

type DeviceResponse struct {
	Imei    string `json:"imei" validate:"required"`
	Capital string `json:"capital" validate:"required"`
	ComName string `json:"computer" validate:"required"`
}
