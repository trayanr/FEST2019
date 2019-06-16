package drivers

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/gchaincl/dotsql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/trayanr/FEST2019/models"
)

func TestSetUp(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/fest2019") // 3306
	if err != nil {
		t.Error(err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}
	dot, err := dotsql.LoadFromFile("../sql/auth.sql")

	res, err := dot.Exec(tx, "create-user-table")

	if err != nil {
		t.Error(err)
		tx.Rollback()
	}
	fmt.Println(res)
}

func TestGetUserByCredentials(t *testing.T) {
	user, err := GetUserByCredentials("trayan", "r")
	fmt.Println(user, err)
	user, err = GetUserByCredentials("t", "r")
	fmt.Println(user, err)
}

func TestInster(t *testing.T) {
	u1 := models.User{Username: "trayan", Password: "r"}
	err := InsertUser(u1)
	fmt.Println(err)
}
