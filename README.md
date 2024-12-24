# Stock-Trading-System-Low-Level-Design-in-GoLang

A robust Go-based stock trading system implementing low-level design patterns for managing stock portfolios, watchlists, and order execution.

## System Overview

This system provides a comprehensive solution for stock trading operations with key features including:
- Portfolio management
- Stock watchlist functionality
- Order execution and validation
- Stock inventory tracking
- User account management with funds handling

![image](https://github.com/user-attachments/assets/aa2b50d0-82fc-49c2-a006-ae17deb67d14)

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

## Command-Line Results

<img width="887" alt="image" src="https://github.com/user-attachments/assets/2e6fb5db-b53c-40a4-ad83-65206f75f17e" />

## Design Patterns Used
- Singleton pattern for OrderManager
- Factory pattern for Order creation
- Observer pattern for Watchlist updates
- Strategy pattern for Order execution

## Error Handling
The system implements comprehensive error handling for:
- Insufficient funds
- Order validation
Here's how you can structure the `README.md` file for your project to guide users on how to clone the repository, navigate to the project directory, and run the Go application:

## Prerequisites

Before running the application, make sure you have the following installed:

- [Go](https://golang.org/doc/install) version 1.18 or above.

## Steps to Run the Application

### 1. Clone the Repository

Start by cloning the project repository to your local machine. Open a terminal and run the following command:

```bash
git clone https://github.com/VishalPrasanna11/Zerodha-Low-Level-Design-in-GoLang.git
```

### 2. Navigate to the Project Directory

Once the repository is cloned, navigate into the project directory:

```bash
cd Zerodha-Low-Level-Design-in-GoLang
```

### 3. Install Go Dependencies

Ensure that your Go modules are set up and any dependencies are installed. Run the following command to download and install any required dependencies:

```bash
go mod tidy
```

### 4. Run the Application

You can now run the application using the following command:

```bash
go run main.go
```
## Future Enhancements
- Real-time price updates
- Multiple exchange support
- Advanced order types
- Historical trade tracking
- Performance analytics
- API integration capabilities

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.
