package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	
	// Reading the files named model.conf and policy.csv
	// Understanding the login
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(err)
	}

	// Checks if the alice can access the data
	ok, err := e.Enforce("alice", "data1", "write")
	if err != nil {
		panic(err)
	}

	if ok {
		fmt.Println("✅ Access granted")
	} else {
		fmt.Println("❌ Access denied")
	}
}
