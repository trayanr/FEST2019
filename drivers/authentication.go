package drivers

import (
	"log"

	"github.com/gchaincl/dotsql"
	"github.com/trayanr/FEST2019/models"

	_ "github.com/go-sql-driver/mysql" // sql driver
)

func GetUserByCredentials(username, password string) (models.User, error) {
	db, err := OpenSQL()
	if err != nil {
		log.Println(errOpenSQl, ":", err)
		return models.User{}, errOpenSQl
	}
	dot, err := dotsql.LoadFromFile("./sql/auth.sql")
	if err != nil {
		log.Println(errLoadDotSQLFile, ":", err)
		return models.User{}, errLoadDotSQLFile
	}
	row, err := dot.QueryRow(db, "get-user-by-username-passwordhash", username, password)
	if err != nil {
		log.Println(errSQL, ":", err)
		return models.User{}, errSQL
	}
	user := models.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		//handle err
		return models.User{}, err
	}
	return user, nil
}

func InsertUser(user models.User) error {
	db, err := OpenSQL()
	if err != nil {
		log.Println(errOpenSQl, ":", err)
		return errOpenSQl
	}
	dot, err := dotsql.LoadFromFile("./sql/auth.sql")
	if err != nil {
		log.Println(errLoadDotSQLFile, ":", err)
		return errLoadDotSQLFile
	}

	_, err = dot.Exec(db, "create-user", user.Username, user.Email, user.Password)
	if err != nil {
		log.Println(errSQL, err)
		return errSQL
	}
	return nil
}
