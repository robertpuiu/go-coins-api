# Go Coins API

A simple web API built with Go, utilizing the Chi router and Logrus for logging. It provides endpoints to manage user accounts and retrieve coin balances.

## Features

- Retrieve coin balances for user accounts.
- Middleware for authorization.
- Mock database implementation for demonstration purposes.

## Project Structure

```plaintext
go-coins-api/
├── cmd/
│   └── api/
│       └── main.go          # Entry point of the application
├── internal/
│   ├── handlers/            # HTTP handlers for API endpoints
│   ├── middleware/          # Authorization middleware
│   └── tools/               # Utilities, including mock database
├── api/
│   └── api.go               # API request and response definitions
├── go.mod                   # Go module file
└── go.sum                   # Go dependencies checksum file
```
## Installation

#### 1. Clone the repository:

```
git clone git@github.com:robertpuiu/go-coins-api.git
```

#### 2. Navigate to the project directory:
```
cd go-coins-api
```

#### 3. Download dependencies
```
go mod download
```

## Running the Application

#### 1.Navigate to the cmd/api directory
```
cd cmd/api
```

#### 2.Run the application:
```
go run main.go
```

#### The API server will start at http://localhost:8000

## Usage

#### Request exemple:
```
curl -H "Authorization: 944PUIU" "http://localhost:8000/account/coins?username=robert"
```

#### Mock Database

Username	Auth Token	Coin Balance
robert	944PUIU	9441
vlad	001STEF	861
alex	004LEAF	1285
Username | Auth Token | Coin Balance 
--- | --- | ---
robert | 944PUIU | 9441 
vlad | 001STEF | 861 
alex | 004LEAF | 1285 
