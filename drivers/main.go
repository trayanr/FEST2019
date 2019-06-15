package drivers

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

var (
	errOpenSQl        = errors.New("could not open connection to sql server")
	errLoadDotSQLFile = errors.New("could not open sql file")
	errSQL            = errors.New("an error wiht the sql occured")
)

func OpenSQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/fest2019")
	return db, err
}
