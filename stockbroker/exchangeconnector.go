package stockbroker

import (
	"sync"
)

type ExchConnector struct{}

var (
	instance *ExchConnector
	once     sync.Once
)

func GetInstance() *ExchConnector {
	once.Do(func() {
		instance = &ExchConnector{}
	})
	return instance
}

func (ec *ExchConnector) PlaceOrder(order *Order) error {
	// Interact with exchange
	return nil
}
