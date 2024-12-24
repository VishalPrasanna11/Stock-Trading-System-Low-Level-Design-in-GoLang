package stockbroker

import "time"

type Order struct {
	orderId    int
	exchange   EXCHANGE_TYPE
	stock      *STOCK
	price      float64
	quantity   int
	orderType  Order_TYPE
	txnType    TXN_TYPE
	status     ORDER_STATUS
	order_time time.Time
	totalPrice float64
}

// NewOrder creates a new Order and calculates total price
func NewOrder(porderId int, pexchange EXCHANGE_TYPE, pstock *STOCK, pprice float64, pquantity int, porderType Order_TYPE, ptxnType TXN_TYPE, pstatus ORDER_STATUS) Order {
	totalPrice := pprice * float64(pquantity)
	return Order{
		orderId:    porderId,
		exchange:   pexchange,
		stock:      pstock,
		price:      pprice,
		quantity:   pquantity,
		orderType:  porderType,
		txnType:    ptxnType,
		status:     pstatus,
		order_time: time.Now(),
		totalPrice: totalPrice,
	}
}

// GetTxnType returns the transaction type of the order
func (o *Order) GetTxnType() TXN_TYPE {
	return o.txnType
}

// GetTotalPrice returns the total price of the order
func (o *Order) GetTotalPrice() float64 {
	return o.totalPrice
}
