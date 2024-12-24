package stockbroker

type OrderManager struct {
	OrderValidator   *OrderValidator
	OrderExecutioner *OrderExecutioner
}

// PlaceOrder validates and executes the order
func (om *OrderManager) PlaceOrder(userId string, order *Order) {
	if om.OrderValidator.ValidateOrder(userId, order) {
		om.OrderExecutioner.ExecuteOrder(order, userId)
	}
}

// NewOrderManager initializes an OrderManager with a given UserList
func NewOrderManager(userList *UserList) *OrderManager {
	return &OrderManager{
		OrderValidator:   NewOrderValidator(userList),           // Use the provided UserList
		OrderExecutioner: &OrderExecutioner{userList: userList}, // Use the provided UserList
	}
}
