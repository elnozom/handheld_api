package model

type GetItemRequest struct {
	BCode     string `json:"BCode" validate:"required"`
	Name      string `json:"Name"`
	StoreCode int    `json:"StoreCode" validate:"required"`
}

type Item struct {
	Serial            int
	ItemName          string
	MinorPerMajor     int
	POSPP             float64
	POSTP             float64
	I                 float64
	R                 float64
	ByWeight          bool
	WithExp           bool
	ItemHasAntherUnit bool
	AvrWait           float64
	Expirey           string
	LastBuyPrice      float64
}
type ItemInInvReq struct {
	ItemSerial int
	BonSerial  int
}
type ItemInInvResp struct {
	Found int8
}
type InsertItemReq struct {
	DNo        int
	TrS        int
	AccS       int
	ItmS       int
	Qnt        float32
	StCode     int
	StCode2    int
	InvNo      int
	ItmBarCode string
	DevNo      int
	ExpDate    string
	SessionNo  int
}

type DocItemsReq struct {
	DocNo     int
	TrSerial  int
	StoreCode int
	DevNo     int
}

type DocItem struct {
	Serial        int
	Qnt           float32
	Price         float64
	TotalCash     float64
	Item_BarCode  int
	MinorPerMajor int
	ItemTotal     float64
	ItemName      string
	ByWeight      bool
}

type DeleteItemReq struct {
	Serial int
}
