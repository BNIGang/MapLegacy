package v1

import "github.com/BNIGang/MapLegacy/web"

type UserList struct {
	User_ID       string
	Name          string
	Username      string
	Wilayah       string
	UserPrivilege string
	CabangName    string
}

var userMap = make(map[string]*UserList)

func GetUsers() ([]UserList, error) {
	userMap = make(map[string]*UserList)

	db, err := web.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			u.user_id, 
			u.name, 
			u.username, 
			up.wilayah_id, 
			up.user_privilege, 
			c.cabang_name 
		FROM 
			users u 
		LEFT JOIN 
			user_privileges up 
		ON 
			u.user_id=up.user_id 
		LEFT JOIN 
			cabang c 
		ON 
			up.cabang_id=c.cabang_id
    `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userList []UserList

	for rows.Next() {
		var userlist UserList

		err = rows.Scan(
			&userlist.User_ID,
			&userlist.Name,
			&userlist.Username,
			&userlist.Wilayah,
			&userlist.UserPrivilege,
			&userlist.CabangName,
		)
		if err != nil {
			return nil, err // database error
		}

		userList = append(userList, userlist)

		// Store the user in the userMap using the User_ID as the key
		userMap[userlist.User_ID] = &userlist
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userList, nil
}
