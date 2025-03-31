package test

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestCasbinMysql(t *testing.T) {

	a, _ := gormadapter.NewAdapter("mysql", "root:123456@tcp(121.199.44.128:3306)/")
	e, err := casbin.NewEnforcer("./rbac_model.conf", a)
	e.LoadPolicy()
	//rules := [][]string{
	//	[]string{"admin", "/group/:id", "GET"},
	//	[]string{"admin", "/group", "POST"},
	//	[]string{"admin", "/group/:id", "PUT"},
	//	[]string{"admin", "/group/:id", "DELETE"},
	//}
	//_, err = e.AddPolicies(rules)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//groups := [][]string{
	//	[]string{"user1", "admin"},
	//	[]string{"user2", "admin"},
	//}
	//e.AddGroupingPolicies(groups)

	//err = e.SavePolicy()
	//if err != nil {
	//	log.Fatal(err)
	//}

	ok, err := e.Enforce("user1", "/goup/123334", "PUT")
	if ok {

		fmt.Println(ok, err)
	}
}
