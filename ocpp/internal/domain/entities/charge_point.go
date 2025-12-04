package entities

type ChargePoint struct {
	ID             string
	Vendor         string
	Model          string
	LastHearthbeat int64
}
