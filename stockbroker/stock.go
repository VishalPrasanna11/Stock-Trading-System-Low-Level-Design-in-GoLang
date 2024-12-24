package stockbroker

type STOCK struct {
	exchange EXCHANGE_TYPE
	name     string
	price    float64
}

func NewStock(pexchange EXCHANGE_TYPE, pname string, pprice float64) STOCK {
	return STOCK{exchange: pexchange, name: pname, price: pprice}
}
