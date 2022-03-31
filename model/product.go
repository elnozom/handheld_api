package model

type ProductCreateInitialReq struct {
	ItemCode           int
	GroupCode          int
	BarCode            string
	Name               string
	MinorPerMajor      int
	AccountSerial      int
	ActiveItem         bool
	ItemTypeID         int
	ItemHaveSerial     bool
	MasterItem         bool
	ItemHaveAntherUint bool
	StoreCode          int
	LastBuyPrice       float64
	POSTP              float64
	POSPP              float64
	Ratio1             float64
	Ratio2             float64
}

type GroupCodeAndMaxItem struct {
	GroupName string
	MaxCode   int
}
type ItemType struct {
	ItemTypeID   int
	ItemTypeName string
}
