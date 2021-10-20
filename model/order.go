package model

type OrderItem struct {
	Serial   int
	BarCode  int
	ItemName string
	Qnt      float64
	Price    float64
	Total    float64
}

type InsertOrder struct {
	DocNo         int
	StoreCode     int
	EmpCode       int
	AccountSerial int
}

type GetOrderItemsRequest struct {
	Serial int `json:"Serial" validate:"required"`
}
type InsertOrderItem struct {
	HeadSerial    int
	ItemSerial    int
	QntAntherUnit float64
	Qnt           float64
	Price         float64
}
