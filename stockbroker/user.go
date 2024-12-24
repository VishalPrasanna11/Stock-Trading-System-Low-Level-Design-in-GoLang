package stockbroker

// User represents a user in the system
type User struct {
	userId    string
	funds     float64
	watchlist *Watchlist
	portfolio *Portfolio
}

// NewUser creates a new user with a watchlist and portfolio
func NewUser(puserId string, pfunds float64) *User {
	return &User{
		userId:    puserId,
		funds:     pfunds,
		watchlist: NewWatchlist(),
		portfolio: NewPortfolio(),
	}
}

// AddFunds adds funds to the user's balance
func (u *User) AddFunds(amount float64) {
	u.funds += amount
}

// DeductFunds deducts funds from the user's balance
func (u *User) DeductFunds(amount float64) bool {
	if u.funds >= amount {
		u.funds -= amount
		return true
	}
	return false
}

// ListStocks lists all stocks in the user's portfolio
func (u *User) ListStocks() []STOCK {
	return u.portfolio.ListStocks()
}

// AddStock adds a stock to the user's portfolio
func (u *User) AddStock(stock STOCK) {
	u.portfolio.AddStock(stock)
}

// AddToWatchlist adds a stock to the user's watchlist
func (u *User) AddToWatchlist(stock STOCK) {
	u.watchlist.AddToWatchlist(stock)
}
