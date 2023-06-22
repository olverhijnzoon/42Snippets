package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	Name   string
	Email  string
	Active bool
}

type Users struct {
	data map[string]*User
	mu   sync.RWMutex
}

func NewUsers() *Users {
	return &Users{
		data: make(map[string]*User),
	}
}

func (u *Users) AddUser(name string, user *User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.data[name] = user
}

func (u *Users) RemoveUser(name string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	delete(u.data, name)
}

func (u *Users) GetUser(name string) (*User, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	user, ok := u.data[name]
	return user, ok
}

func main() {
	fmt.Println("# 42Snippets")
	fmt.Println("## Mutex")

	/*
	   This code exemplifies the use of a mutex (short for "mutual exclusion") to ensure thread safety during concurrent operations on a shared resource, in this case a map of User pointers. The mutex is part of a Users struct which includes the shared map and methods to AddUser, RemoveUser, and GetUser from the map.

	   To illustrate the effectiveness of the mutex, two goroutines are spawned - one adding a user to the map and one removing the user. Even though these operations are called concurrently, the mutex prevents simultaneous access and modification of the map, thus averting any race condition.

	   The `rand.Intn` function and `time.Sleep` are used to introduce an unpredictable delay before each operation, making it ambiguous when each goroutine will attempt to lock the mutex and access the map. Regardless of when the attempt is made, the mutex ensures that one goroutine's operation will not interrupt the other's.

	   A `sync.WaitGroup` is employed to make the main function wait until both goroutines finish their operations before it attempts to retrieve the user from the map.

	   The end result of the program is either "Michael Schmidt" or "User not found in map", depending on the sequence of lock acquisition and release between the AddUser and RemoveUser goroutines.
	*/

	users := NewUsers()
	var wg sync.WaitGroup
	wg.Add(2)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Add a user
	go func() {
		defer wg.Done()
		u := &User{
			Name:   "Michael Schmidt",
			Email:  "michael@marslabs.mars",
			Active: true,
		}
		users.AddUser("michael", u)
	}()

	// Random sleep before removing the user
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	// Remove a user
	go func() {
		defer wg.Done()
		users.RemoveUser("michael")
	}()

	wg.Wait() // Wait for both goroutines to finish

	u, ok := users.GetUser("michael")
	if !ok {
		fmt.Println("User not found in map")
	} else {
		fmt.Println(u.Name)
	}
}
