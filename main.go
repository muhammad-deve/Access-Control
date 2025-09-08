package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	// Create a new Enforcer by loading:
	// - model.conf → defines the logic (how requests & policies are matched)
	// - policy.csv → defines the rules (who can do what)
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(err)
	}

	// Define the request we want to check:
	// "alice" = subject (user)
	// "data1" = object (resource)
	// "write" = action
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
