package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	// Load model and policy
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(err)
	}

	// Check: alice -> data1 -> write (since alice is in role admin)
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
