package stockbroker

// Watchlist holds the stocks a user is interested in
type Watchlist struct {
	stocks map[string]STOCK
}

// NewWatchlist creates a new empty watchlist
func NewWatchlist() *Watchlist {
	return &Watchlist{
		stocks: make(map[string]STOCK),
	}
}

// AddToWatchlist adds a STOCK to the watchlist
func (w *Watchlist) AddToWatchlist(stock STOCK) {
	w.stocks[stock.name] = stock
}

// RemoveFromWatchlist removes a stock from the watchlist
func (w *Watchlist) RemoveFromWatchlist(stockName string) {
	delete(w.stocks, stockName)
}

// ListWatchlist lists all stocks in the watchlist
func (w *Watchlist) ListWatchlist() []STOCK {
	var stockList []STOCK
	for _, stock := range w.stocks {
		stockList = append(stockList, stock)
	}
	return stockList
}
