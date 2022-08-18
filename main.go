package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	// Initialize a Xorm adapter with MySQL database.
	a, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}
	m, _ := model.NewModelFromFile("conf/basic_model.conf")

	e, err := casbin.NewEnforcer(m, a)

	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	fmt.Println(e)
}
