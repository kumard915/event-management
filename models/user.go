package models

import (
	"errors"
	"fmt"

	"example.com/restapi/db"
	"example.com/restapi/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	pass, err := utils.HashedPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, pass)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userID
	return nil
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)
	var retrivePassword string

	err := row.Scan(&u.ID, &retrivePassword)
	fmt.Println("Error:", err, "Retrieved Password:", retrivePassword)
	if err != nil {
		return errors.New("credential invalid")
	}

	passwordIsValid := utils.CheckPasswordHashed(u.Password, retrivePassword)

	if !passwordIsValid {
		return errors.New("credential invalid")
	}

	return nil
}
