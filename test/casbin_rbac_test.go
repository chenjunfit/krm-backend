package test

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

func check() bool {
	return true
}
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	return (bool)(check()), nil
}

func TestCasbinRBAC(t *testing.T) {

	e, _ := casbin.NewEnforcer("./rbac_model.conf", "./rbac_policy.csv")
	e.LoadPolicy()
	e.AddFunction("check", KeyMatchFunc)
	ok, _ := e.Enforce("admin", "/group/234", "DELETE")
	fmt.Println(ok)
	if ok {
		fmt.Println("admin can post /group")
	}

}
