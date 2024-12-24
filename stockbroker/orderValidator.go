package stockbroker

import (
	"fmt"
)

// OrderValidator handles validation of user trading orders
type OrderValidator struct {
	userList *UserList
}

// NewOrderValidator creates a new instance of OrderValidator
func NewOrderValidator(ul *UserList) *OrderValidator {
	return &OrderValidator{userList: ul}
}

// ValidateOrder checks if a user can place the given order
func (v *OrderValidator) ValidateOrder(userID string, order *Order) bool {

	user := v.userList.GetUser(userID)
	if user == nil {
		fmt.Printf("This is from the order validator: User %s not found\n", userID)
		return false
	}

	switch order.GetTxnType() {
	case BUY:
		fmt.Printf("Checking if user %s has enough funds to buy\n", userID)
		if user.funds < order.totalPrice {
			fmt.Printf("Insufficient funds: User %s has %.2f, but order requires %.2f\n", userID, user.funds, order.totalPrice)
			return false
		}
		fmt.Println("Funds are sufficient for the BUY transaction.")
	case SELL:
		fmt.Printf("Checking if user %s has enough stocks to sell in portfolio\n", userID)
		// Assuming a portfolio check for SELL is required here
		fmt.Println("For now, let's assume user has enough stocks and return true")
	default:
		fmt.Printf("Unknown transaction type for user %s\n", userID)
		return false
	}

	return true
}
