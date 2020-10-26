package utils

import (
	"github.com/notify/genaral/models"
	"github.com/notify/db"
_ "github.com/go-sql-driver/mysql"
)


//AddUser add user to data base
func InsertUser(User model.UserInfo) error {
	
	dbConnection, err := db.GetMySQLConnection()
	if(err !=nil){
		return err
	}

	stmtIns, err := dbConnection.Prepare("INSERT INTO users(first_name, last_name, user_name, password, role) VALUES( ?, ?,?,?,?,? )") // ? = placeholder

	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(User.FirstName,User.LastName, User.Email, User.Password, User.Role)

	if err != nil {
		return err
	}

	stmtIns.Close()

	return nil

}

//FindUser add user to data base
func GetUserInfo(Email string)(*model.UserInfo , error)  {
	userInfo := &model.UserInfo{};
	dbConnection, err := db.GetMySQLConnection()
	if(err !=nil){
		return nil, err
	}
	stmtOut, err := dbConnection.Prepare("SELECT user_id, first_name, last_name, user_name, password, role, is_active FROM users WHERE user_name	 = ? AND is_active = 1")
	if err != nil {
		return nil, err
	}


	err = stmtOut.QueryRow(Email).Scan(
		&userInfo.UserId,
		&userInfo.FirstName,
		&userInfo.LastName,
		&userInfo.Email ,
		&userInfo.Password,
		&userInfo.Role,
		&userInfo.IsActive, 
	)
	if err != nil {
		return nil, err
	}
	return userInfo, err

}
