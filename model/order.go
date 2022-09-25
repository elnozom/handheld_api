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

type ApplyDiscount struct {
	Discount float64
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
	DocDate       string
	Discount      float64
	StoreCode     int
	DocNo         string
	AccountSerial int
	TotalCash     float64
	TransSerial   int
	AccountName   string
	AccountCode   int
}

type PrintItem struct {
	ItemName   string
	WholeQnt   float64
	PartQnt    float64
	Price      float64
	Tax        float64
	TotalPrice float64
	GrandTotal float64
	BarCode    string
}
type PrintHeader struct {
	DocDate        string
	Discount       float64
	IsClosed       bool
	AccountName    string
	AccountTrn     string
	AccountAddress string
	DocNo          string
	EmpName        string
	StoreName      string
	Total          float64
	Tax            float64
	SubTotal       float64
	GrandTotal     float64
}
type PrintResponse struct {
	Items  []PrintItem
	Header PrintHeader
	Info   CompanyInfo
}
type ListDirectOrdersReq struct {
	StoreCode    int     `query:"store"`
	TransSerial  int     `query:"trSerial"`
	IsClosed     bool    `query:"isClosed"`
	FromDate     *string `query:"fromDate"`
	ToDate       *string `query:"toDate"`
	ComputerName string  `query:"computer"`
}
