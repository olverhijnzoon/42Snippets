package main

import (
	"fmt"
	"math/rand"
	"time"

	mystore "github.com/olverhijnzoon/42Snippets/golanginterfaces/mystore"
	storacle "github.com/olverhijnzoon/42Snippets/golanginterfaces/storacle"
)

/*
	"Declare an interface where it is used, not where it is implemented. Unless the interface is well discovered." https://twitter.com/inancgumus/status/1605116543553019905?s=20&t=fTo98MYIjsZ5nw44XnTmRw

	In Go, it is generally considered good practice to declare an interface in the same package where it is used, rather than where it is implemented. This is because interfaces are a way of specifying the behavior of a type, rather than the implementation of that type. By declaring an interface in the same package where it is used, you make it clear to users of the interface what they can expect from types that implement it.

	For example, suppose we have a package called mystore that defines an interface called Store for storing and retrieving data. If we declare the Store interface in the mystore package, users of the mystore package can see exactly what methods they can call on types that implement the Store interface.

	On the other hand, if we declare the Store interface in a separate package where it is implemented, users of the mystore package would not be able to see the methods of the Store interface unless they imported the implementation package. This could make it more difficult for users of the mystore package to understand how to use types that implement the Store interface.

	There are, however, cases where it may be appropriate to declare an interface in a separate package where it is implemented. For example, if the interface is intended to be a "well-known" interface that will be used by many packages, it may make sense to declare it in a separate package to make it easier for other packages to import and use.

	In the following example, the Store interface is declared here in the main package and specifies two methods: Save and Load. Any type that implements these methods can be used as a Store.
*/

// Store is an interface for storing and retrieving data.
type Store interface {
	// Save saves the given data to the store.
	Save(data []byte) error

	// Load loads the data from the store.
	Load() ([]byte, error)
}

func main() {

	// Create a variable of type Store
	var s Store

	// Set s to a value that implements the Store interface like MyStore from mystore package
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	if rand.Intn(2) == 0 {           // Generate a random number that is either 0 or 1
		s = &mystore.MyStore{}
	} else {
		s = &storacle.MyStore{}
	}

	// Call a function that expects a Store
	err := processData(s)
	if err != nil {
		// Handle the error...
	}
}

func processData(s Store) error {
	data, err := s.Load()
	if err != nil {
		return err
	}

	// Process the data...
	fmt.Println("Processing the data!")

	return s.Save(data)
}
