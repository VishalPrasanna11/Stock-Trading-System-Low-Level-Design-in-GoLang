package stockbroker

import "fmt"

// Run encapsulates the main application logic
func Run() {
	// Create a new user
	user := NewUser("1", 100000.0)

	// Create a new user list and add the user to the list
	userList := NewUserList()
	userList.AddUser(user.userId, user)

	// Create a new stock inventory
	stockInventory := NewStockInventory()

	// Add stocks to the inventory (e.g., NVIDIA, Apple)
	stockInventory.AddStock(NewStock("NSE", "NVIDIA", 1000.0))
	stockInventory.AddStock(NewStock("NSE", "APPLE", 150.0))

	// Print all stocks in the inventory
	fmt.Println("Stock Inventory:")
	for _, stock := range stockInventory.ListStocks() {
		fmt.Println(stock.name)
	}

	// User adds stock to the watchlist
	user.AddToWatchlist(NewStock("NSE", "NVIDIA", 1000.0))

	// List the watchlist
	fmt.Println("\nUser's Watchlist:")
	for _, stock := range user.watchlist.ListWatchlist() {
		fmt.Println(stock.name)
	}

	// Create a new order manager and pass the user list
	orderManager := NewOrderManager(userList)

	// Create a new order to buy stock (e.g., Buy 5 NVIDIA shares)
	stock, exists := stockInventory.GetStock("NVIDIA")
	if !exists {
		fmt.Println("Stock not found")
		return
	}
	order := NewOrder(1, "NSE", &stock, 1000.0, 5, "LIMIT", "BUY", "OPEN")

	// Place the order for the user using the OrderManager
	orderManager.PlaceOrder(user.userId, &order)

	// Print the portfolio after the stock has been added
	fmt.Println("\nUser's Portfolio:")
	for _, stock := range user.ListStocks() {
		fmt.Println(stock.name)
	}
}
