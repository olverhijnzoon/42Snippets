package main

import "fmt"

// Define a struct type named 'Repository'
type Repository struct {
	Name        string
	PublicLink  string
	Description string
	Language    string
	Stars       int
	Forks       int
	LastUpdated string
}

// Define another struct type named 'Company'
type Company struct {
	Name, Category       string
	MarketCapitalization float64 // In billions USD
	Repository
}

func main() {
	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Structs")

	/*
		In Go, custom data types are represented using structs. One powerful feature of Go's struct types is that they support embedding, which allows one struct to be included as a field within another struct.
		TODO: Struct Tags
	*/

	// Initialize new 'Repository' struct
	deepLearningExamples := Repository{
		Name:        "DeepLearningExamples",
		PublicLink:  "https://github.com/NVIDIA/DeepLearningExamples",
		Description: "Deep Learning Examples for NVIDIA Tensor Cores",
		Language:    "Python",
		Stars:       3500,         // Hypothetical
		Forks:       1000,         // Hypothetical
		LastUpdated: "2023-05-31", // Hypothetical
	}

	// Initialize new 'Company' structs
	tesla := Company{
		Name:                 "Tesla",
		Category:             "AI",
		MarketCapitalization: 650.0,
	}

	nvidia := Company{
		Name:                 "NVIDIA",
		MarketCapitalization: 500.0,
		Repository:           deepLearningExamples,
	}

	// Access struct fields
	fmt.Printf("Company: %s, Category: %s, Market Cap: $%.0fB\n", tesla.Name, tesla.Category, tesla.MarketCapitalization)    //
	fmt.Printf("Company: %s, Category: %s, Market Cap: $%.0fB\n", nvidia.Name, nvidia.Category, nvidia.MarketCapitalization) //
	fmt.Printf("Name: %s\nLink: %s\nDescription: %s\nLanguage: %s\nStars: %d\nForks: %d\nLast Updated: %s\n", nvidia.Repository.Name, nvidia.Repository.PublicLink, nvidia.Repository.Description, nvidia.Repository.Language, nvidia.Repository.Stars, nvidia.Repository.Forks, nvidia.Repository.LastUpdated)
}
