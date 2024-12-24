package stockbroker

import "fmt"

// OrderExecutioner is responsible for sending orders to ExchConnector
type OrderExecutioner struct {
	userList *UserList // Ensure userList is accessible in the OrderExecutioner struct
}

// ExecuteOrder sends the order to the ExchConnector singleton instance and updates the user's portfolio
func (oe *OrderExecutioner) ExecuteOrder(order *Order, userID string) {
	// Get the instance of the exchange connector
	exchConnector := GetInstance()
	err := exchConnector.PlaceOrder(order)
	if err != nil {
		fmt.Printf("Failed to place order: %v\n", err)
		return
	}

	// Retrieve the user from the user list
	user := oe.userList.GetUser(userID)
	if user == nil {
		fmt.Println("User not found")
		return
	}

	fmt.Println("Order placed successfully!")

	// Add the stock to the user's portfolio after a successful order
	user.AddStock(*order.stock)
}
