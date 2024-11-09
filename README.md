# Go - DDD Onion Project Structure

This project is a simple example of how to structure a Go project using the DDD (Domain-Driven Design) and Onion Architecture concepts.

The project is a simple API that allows you to handle orders.

The project is used as reference for my talk about DDD and Onion Architecture in Go.

Any suggestions or improvements are welcome.

## Project Structure

```
.
├── pkg
│   ├── order // Order domain
│   │   ├── infra // Infrastructure layer
│   │   ├── repos // Repository implementation
│   │   ├── service // Application layer
│   │   └── order.go // Order domain
│   │   └── order_repo.go // Order repository interface
|   └── ...
├── internal
│   └── errors // Custom errors
├── cmd // Application entry point
├── main.go // Main file
└── go.mod // Go modules file
```
