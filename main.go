package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

// 🔑 Casbin Terminology
// Subject (sub) → The actor making the request (e.g., user, admin, service)
// Object  (obj) → The resource being accessed (e.g., data, API, file)
// Action  (act) → The operation performed on the resource (e.g., read, write, delete)

func main() {

	// Reading the files named model.conf and policy.csv
	// Understanding the logic
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(err)
	}

	// ✅ You must load the policies
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	// Checks if the subject can access the data
	ok, err := e.Enforce("bob", "data1", "read")
	if err != nil {
		panic(err)
	}

	if ok {
		fmt.Println("✅ Access granted")
	} else {
		fmt.Println("❌ Access denied")
	}
}
