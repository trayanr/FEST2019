package drivers

import (
	"errors"
	"fmt"
	"log"

	"github.com/kachamaka/argon2custom"

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
	row, err := dot.Query(db, "get-user-by-username", username)
	if err != nil {
		log.Println(errSQL, ":", err)
		return models.User{}, errSQL
	}
	// users := []models.User{}
	for row.Next() {
		user := models.User{}
		err = row.Scan(&user.ID, &user.Username, &user.Password, &user.OAuthCode, &user.Level, &user.Points, &user.LastChecked)
		fmt.Println("--PASS--")
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(user.Password)
		ok, err := argon2custom.ComparePasswordAndHash(password, user.Password)
		fmt.Println(ok, err)
		if ok {
			return user, nil
		}
	}
	if err != nil {
		//handle err
		return models.User{}, err
	}
	return models.User{}, errors.New("no such user")
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

	_, err = dot.Exec(db, "create-user", user.Username, user.Password)
	if err != nil {
		log.Println(errSQL, err)
		return errSQL
	}
	return nil
}

func SetUserAuthCode(authCode string, id int) error {
	db, err := OpenSQL()
	if err != nil {
		return errOpenSQl
	}
	dot, err := dotsql.LoadFromFile("./sql/auth.sql")
	if err != nil {
		return errLoadDotSQLFile
	}
	_, err = dot.Exec(db, "set-user-auth", authCode, id)
	if err != nil {
		return errSQL
	}
	return nil
}

func GetUserByID(id int) (models.User, error) {
	db, err := OpenSQL()
	if err != nil {
		return models.User{}, errOpenSQl
	}
	dot, err := dotsql.LoadFromFile("./sql/auth.sql")
	if err != nil {
		return models.User{}, errLoadDotSQLFile
	}
	rows, err := dot.Query(db, "get-user-by-id", id)
	user := models.User{}
	fmt.Println(rows)
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.OAuthCode, &user.Level, &user.Points, &user.LastChecked)
		fmt.Println(user, id, err)
	}
	sessions, err := models.GetLastSession(user.OAuthCode, int64(user.LastChecked))
	fmt.Println(err, sessions)

	points := models.Points{
		Value: user.Points,
		Level: user.Level,
	}
	points.Calculate(sessions[0].ActivityType, uint(sessions[0].EndTimeMillis-sessions[0].StartTimeMillis))
	fmt.Println(points.Value)
	_, err = dot.Exec(db, "update-points", points.Value, points.Level, user.ID)
	fmt.Println(err)
	_, err = dot.Exec(db, "set-lastChecked", sessions[0].EndTimeMillis, user.ID)
	fmt.Println(err)
	return user, nil
}
