package test

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

func TestCasbinRBAC(t *testing.T) {

	e, _ := casbin.NewEnforcer("./rbac_model.conf", "./rbac_policy.csv")

	ok, _ := e.Enforce("test", "/group", "POST")
	if ok {
		fmt.Println("admin can post /group")
	}

}
