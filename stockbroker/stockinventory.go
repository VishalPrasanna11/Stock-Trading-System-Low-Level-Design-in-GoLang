package stockbroker

type StockInventory struct {
	stocks map[string]STOCK
}

// NewStockInventory creates a new stock inventory
func NewStockInventory() *StockInventory {
	return &StockInventory{
		stocks: make(map[string]STOCK),
	}
}

// AddStock adds a stock to the inventory
func (si *StockInventory) AddStock(stock STOCK) {
	si.stocks[stock.name] = stock
}

// GetStock retrieves a stock by name
func (si *StockInventory) GetStock(name string) (STOCK, bool) {
	stock, exists := si.stocks[name]
	return stock, exists
}

// UpdateStock updates the details of an existing stock
func (si *StockInventory) UpdateStock(name string, newPrice float64) bool {
	stock, exists := si.stocks[name]
	if exists {
		stock.price = newPrice
		si.stocks[name] = stock
		return true
	}
	return false
}

// ListStocks lists all stocks in the inventory
func (si *StockInventory) ListStocks() []STOCK {
	var stockList []STOCK
	for _, stock := range si.stocks {
		stockList = append(stockList, stock)
	}
	return stockList
}
