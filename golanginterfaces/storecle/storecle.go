package storacle

import "fmt"

type MyStore struct {
	// Add any fields needed to store data
}

func (s *MyStore) Save(data []byte) error {
	// Implement the Save method
	fmt.Println("Saving with Storecle")
	// Save the data to some storage location (e.g., a database or file)
	return nil
}

func (s *MyStore) Load() ([]byte, error) {
	// Implement the Load method
	fmt.Println("Loading with Storacle")
	// Load the data from the storage location
	return []byte{}, nil
}
