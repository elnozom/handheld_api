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

type InsertDirectOrderReq struct {
	AccountSerial int
	RaseedBefore  float64
	EmpCode       int
	StoreCode     int
	StoreCode2    int
	ComputerName  string
	HeadSerial    int
	TransSerial   int
	ItemSerial    int
	Qnt           float64
	Price         float64
	Tax           float64
	MinorPerMajor int
}

type DirectOrder struct {
	Serial        int
	StoreCode     int
	DocNo         string
	DocDate       string
	Discount      float64
	Total         float64
	AccountSerial int
	TotalCash     float64
	TransSerial   int
	AccountName   string
	AccountCode   int
}

type DirectOrderPrint struct {
	DocDate       string
	DocTime       string
	AccountName   string
	DocNo         string
	ItemName      string
	EmpName       string
	StoreName     string
	WholeQnt      float64
	PartQnt       float64
	MinorPerMajor int
	Price         float64
	TotalPrice    float64
	GrandTotal    float64
}
type PrintTotals struct {
	GrandTotal float64
	WholeQnt   float64
	PartQnt    float64
}
type PrintResponse struct {
	Items  []DirectOrderPrint
	Info   CompanyInfo
	Totals PrintTotals
}
type ListDirectOrdersReq struct {
	StoreCode    int     `query:"store"`
	TransSerial  int     `query:"trSerial"`
	ComputerName string  `query:"computer"`
	IsClosed     bool    `query:"IsClosed"`
	FromDate     *string `query:"FromDate"`
	ToDate       *string `query:"ToDate"`
}
