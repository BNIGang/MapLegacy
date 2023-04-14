package web

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "derpen:@tcp(127.0.0.1:3306)/bni_map_legacy")

	if err != nil {
		return nil, err
	}

	// Check if the connection is working by pinging the database
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GetUserByUsername(username string) (*User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT username, password FROM users WHERE username = ?", username)

	var user User
	err = row.Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // user not found
		}
		return nil, err // database error
	}

	return &user, nil
}

func AuthenticateUser(username string, password string) bool {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false
	}
	if user == nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	}

	return true
}
