package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

type User struct {
	ID    int
	Email string
}

// LogValue implements slog.LogValuer for User.
func (u User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("id", u.ID),
		slog.String("email", u.Email),
	)
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (p Product) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("id", p.ID),
		slog.String("name", p.Name),
		slog.Float64("price", p.Price),
	)
}

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Logging")

	/*
		This Go code defines two data structures: User and Product, each with their respective fields and methods to log their values using the slog logging library. The main function initializes sample data for a user and a product, and then logs two events: adding a product to the cart and completing a purchase, both with associated user, product, and timestamp details. The logging methods utilize the custom LogValue methods of the User and Product types to format their data for logging.

		HandlerOptions contains logging options; AddSource adds the source code position to the log output.
	*/

	// Sample data
	user := User{ID: 123, Email: "42@snippet.s"}
	product := Product{ID: 456, Name: "Laptop", Price: 999.99}

	// User adds a product to the cart
	slog.Info("Product added to cart",
		"user", user,
		"product", product,
		"timestamp", time.Now(),
	)

	// User completes the purchase
	slog.Info("Purchase completed",
		"user", user,
		"product", product,
		"timestamp", time.Now(),
	)

	// Create a logger with JSONHandler
	handlerOptions := slog.HandlerOptions{AddSource: true}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &handlerOptions))
	fmt.Println("---------------------------------------")

	// User adds a product to the cart
	logger.Info("Product added to cart",
		"user", user,
		"product", product,
		"timestamp", time.Now(),
		"env_user", os.Getenv("USER"),
	)

	// User completes the purchase
	logger.Info("Purchase completed",
		"user", user,
		"product", product,
		"timestamp", time.Now(),
		"env_user", os.Getenv("USER"),
	)

	slog.LogAttrs(context.Background(), slog.LevelInfo, "Bye 42Snippets")
}
