package stockbroker

// Portfolio represents a user's collection of stocks
type Portfolio struct {
	stocks map[string]STOCK // Portfolio of stocks owned by the user
}

// NewPortfolio creates a new portfolio
func NewPortfolio() *Portfolio {
	return &Portfolio{
		stocks: make(map[string]STOCK),
	}
}

// AddStock adds a stock to the portfolio
func (p *Portfolio) AddStock(stock STOCK) {
	p.stocks[stock.name] = stock
}

// UpdateStockQuantity updates the quantity of a stock in the portfolio

// ListStocks lists all stocks in the portfolio
func (p *Portfolio) ListStocks() []STOCK {
	var stockList []STOCK
	for _, stock := range p.stocks {
		stockList = append(stockList, stock)
	}
	return stockList
}
