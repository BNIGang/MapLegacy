package web

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Connect() (*sql.DB, error) {
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

type User struct {
	User_ID         string
	Username        string
	Password        string
	Wilayah_ID      string
	Cabang_ID       string
	User_Privileges string
}

func GetUserByUsername(username string) (*User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(`
        SELECT users.user_id, users.username, users.password,
               user_privileges.wilayah_id, user_privileges.cabang_id, user_privileges.user_privilege
        FROM users
        JOIN user_privileges ON users.user_id = user_privileges.user_id
        WHERE users.username = ?`,
		username,
	)

	var user User
	err = row.Scan(&user.User_ID, &user.Username, &user.Password, &user.Wilayah_ID, &user.Cabang_ID, &user.User_Privileges)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // user not found
		}
		return nil, err // database error
	}

	return &user, nil
}

func AuthenticateUser(username string, password string) (bool, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
