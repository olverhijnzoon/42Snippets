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
	Repository           *Repository
}

func updateStars(companyPointer *Company, newStars int) {
	// (*companyPointer).Repository.Stars = newStars // Explicit dereferencing
	companyPointer.Repository.Stars = newStars
}

func newCompany(name, category string, marketcapitalization float64, repository *Repository) *Company {
	return &Company{name, category, marketcapitalization, repository}
}

func main() {
	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Structs")

	/*
		In Go, custom data types are represented using structs. One powerful feature of Go's struct types is that they support embedding, which allows one struct to be included as a field within another struct.

		When you pass a value to a function in Go, the function gets a copy of the value. This is because Go uses a mechanism called "pass by value" where the values of the arguments which are passed in the function call are copied to the parameters of the function. This means if you change the copy inside the function, it does not affect the original value. In Go, you don't necessarily have to use the * operator to access the fields of a struct through a pointer to that struct.

		fmt.Println((*tesla).Repository.Stars) // Explicit dereferencing not necessary

		Constructor function is named "newCompany" by convention.

		TODO: Struct Tags
	*/

	// Initialize new 'Repository' struct
	deepLearningExamples := &Repository{
		Name:        "DeepLearningExamples",
		PublicLink:  "https://github.com/NVIDIA/DeepLearningExamples",
		Description: "Deep Learning Examples for NVIDIA Tensor Cores",
		Language:    "Python",
		Stars:       3500,         // Hypothetical
		Forks:       1000,         // Hypothetical
		LastUpdated: "2023-05-31", // Hypothetical
	}

	// Initialize new 'Company' structs
	tesla := &Company{
		Name:                 "Tesla",
		Category:             "AI",
		MarketCapitalization: 650.0,
		Repository:           &Repository{"Test", "", "", "", 0, 0, ""},
	}

	updateStars(tesla, 6000)

	nvidia := &Company{
		Name:                 "NVIDIA",
		MarketCapitalization: 500.0,
		Repository:           deepLearningExamples,
	}

	// Access struct fields
	fmt.Printf("Company: %s, Category: %s, Market Cap: $%.0fB\n", tesla.Name, tesla.Category, tesla.MarketCapitalization) //
	fmt.Printf("Name: %s\nLink: %s\nDescription: %s\nLanguage: %s\nStars: %d\nForks: %d\nLast Updated: %s\n", (*tesla).Repository.Name, tesla.Repository.PublicLink, tesla.Repository.Description, tesla.Repository.Language, tesla.Repository.Stars, tesla.Repository.Forks, tesla.Repository.LastUpdated)
	fmt.Printf("Company: %s, Category: %s, Market Cap: $%.0fB\n", nvidia.Name, nvidia.Category, nvidia.MarketCapitalization) //
	fmt.Printf("Name: %s\nLink: %s\nDescription: %s\nLanguage: %s\nStars: %d\nForks: %d\nLast Updated: %s\n", nvidia.Repository.Name, (*nvidia).Repository.PublicLink, nvidia.Repository.Description, nvidia.Repository.Language, nvidia.Repository.Stars, nvidia.Repository.Forks, nvidia.Repository.LastUpdated)

	companies := [3]*Company{
		newCompany("Dirac", "Fictional", 0, deepLearningExamples),
		newCompany("Euler", "Fictional", 0, deepLearningExamples),
		newCompany("X", "Fictional", 0, tesla.Repository),
	}
	for _, company := range companies {
		fmt.Println("Name:", company.Name, "Category:", company.Category, "Market Capitalization:", company.MarketCapitalization, "Repository Name:", company.Repository.Name)
	}
}
