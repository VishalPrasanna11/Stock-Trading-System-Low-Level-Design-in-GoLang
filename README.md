# Zerodha-Low-Level-Design-in-GoLang

A robust Go-based stock trading system implementing low-level design patterns for managing stock portfolios, watchlists, and order execution.

## System Overview

This system provides a comprehensive solution for stock trading operations with key features including:
- Portfolio management
- Stock watchlist functionality
- Order execution and validation
- Stock inventory tracking
- User account management with funds handling

## Core Components

### OrderManager
Central component responsible for coordinating order placement and execution. It utilizes:
- `OrderValidator` for validating trade orders
- `OrderExecutioner` for executing validated orders
- Ensures proper order flow from validation to execution

### User Management
Users are managed through the `UserList` component which:
- Maintains a map of user records
- Handles user registration and retrieval
- Manages user-specific data including:
  - Funds (float64)
  - Personal portfolio
  - Custom watchlists
  - Unique user identification

### Portfolio System
Each user has a personal portfolio that:
- Tracks owned stocks using a map[string]STOCK
- Provides methods for adding new stocks
- Lists current stock holdings
- Updates stock quantities after trades

### Watchlist Features
Users can maintain watchlists that:
- Track stocks of interest
- Add/remove stocks dynamically
- List all watched stocks
- Monitor specific stocks for potential trading opportunities

### Order Management
Orders are handled with comprehensive details including:
- Order ID tracking
- Exchange information
- Stock details (name, quantity, price)
- Order type (Buy/Sell)
- Transaction type
- Status tracking
- Timestamp recording
- Total price calculations

### Stock Inventory
The system maintains a central stock inventory that:
- Tracks available stocks
- Manages stock prices
- Updates stock information
- Provides stock availability checks

## Technical Details

### Data Types
- User IDs: string
- Funds: float64
- Stock prices: float64
- Order IDs: int
- Stock quantities: int

### Key Methods

#### User Operations
- `AddFunds(amount float64)`
- `DeductFunds(amount float64) bool`
- `AddStock(stock STOCK)`
- `ListStocks() []STOCK`

#### Order Processing
- `ValidateOrder(userId string, order Order) bool`
- `ExecuteOrder(order Order, userId string)`
- `PlaceOrder(userId string, order Order)`

#### Portfolio Management
- `AddStock(stock STOCK)`
- `ListStocks() []STOCK`

#### Watchlist Management
- `AddToWatchlist(stock STOCK)`
- `RemoveFromWatchlist(stockName string)`
- `ListWatchlist() []STOCK`

## Usage Example

```go
// Create a new user
userList.AddUser("user123", newUser)

// Add funds to user account
user.AddFunds(10000.0)

// Place a buy order
order := Order{
    orderId: 1,
    stock: stock,
    quantity: 10,
    orderType: BUY,
    price: 100.0,
}
orderManager.PlaceOrder("user123", order)
```

## Design Patterns Used
- Singleton pattern for OrderManager
- Factory pattern for Order creation
- Observer pattern for Watchlist updates
- Strategy pattern for Order execution

## Error Handling
The system implements comprehensive error handling for:
- Insufficient funds
- Invalid orders
- Stock availability
- User authentication
- Order validation

## Future Enhancements
- Real-time price updates
- Multiple exchange support
- Advanced order types
- Historical trade tracking
- Performance analytics
- API integration capabilities

## Contributing
Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.