package test

import (
	"fmt"
	"testing"
)
import "github.com/casbin/casbin/v2"

func TestCasbinAcl(t *testing.T) {
	e, err := casbin.NewEnforcer("./basic_model.conf", "./basic_policy.csv")
	if err != nil {
		panic(err)
	}
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
	}
	if ok == true {
		fmt.Println("alice can read data1")
	} else {
		fmt.Println("alice can't read data1")
	}
}
func TestCasbinRbac(t *testing.T) {
	e, err := casbin.NewEnforcer("./rbac_model.conf", "./rbac_policy.csv")
	if err != nil {
		panic(err)
	}
	sub := "alice" // the user that wants to access a resource.
	obj := "data2" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
	}
	if ok == true {
		fmt.Println("alice can read data2")
	} else {
		fmt.Println("alice can't read data2")
	}
}
