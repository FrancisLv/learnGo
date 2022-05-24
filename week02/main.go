package week02

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"week02/dao"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("open db failed")
	}
	defer db.Close()
	queryId := "123456"
	name, err := dao.QueryNameById(db, queryId)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			fmt.Printf("no record was found")
		} else {
			fmt.Printf("query error:\n%+v\n", err)
		}
	}
	fmt.Printf("user name is:%s", name)

}
