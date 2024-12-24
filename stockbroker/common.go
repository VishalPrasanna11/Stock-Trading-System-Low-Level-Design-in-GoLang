package stockbroker

type TXN_TYPE string

const (
	BUY  TXN_TYPE = "BUY"
	SELL TXN_TYPE = "SELL"
)

type Order_TYPE string

const (
	LIMIT  Order_TYPE = "LIMIT"
	MARKET Order_TYPE = "MARKET"
)

type EXCHANGE_TYPE string

const (
	NSE EXCHANGE_TYPE = "NSE"
	BSE EXCHANGE_TYPE = "BSE"
)

type ORDER_STATUS string

const (
	OPEN    ORDER_STATUS = "OPEN"
	PARTIAL ORDER_STATUS = "PARTIAL"
	DONE    ORDER_STATUS = "DONE"
	CANCEL  ORDER_STATUS = "CANCEL"
)
