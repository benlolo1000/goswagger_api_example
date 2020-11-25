package loginUser

import (
	"log"
	"medicarePartD/bus/db/userDatabase"
	"medicarePartD/bus/passwords"
	"medicarePartD/gen/models"
	"regexp"

	"errors"
)

func LoginUser(Email *string, Password *string) error {

	//create database object
	userDb, err := userDatabase.DatabaseCreation()
	if err != nil {
		return err
	}

	// validate email structure
	var validateEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(*Email) > 254 || !validateEmail.MatchString(*Email) {
		return errors.New("not acceptable")
	}

	//Query Email
	rows, err := userDb.Query("SELECT Email, Pass FROM Users WHERE Email=?", *Email)
	if err != nil {
		log.Println(err.Error())
		return errors.New("internal server error")
	}

	//check password is correct
	for rows.Next() {
		var usersList models.Login
		err = rows.Scan(&usersList.Email, &usersList.Password)
		if err != nil {
			log.Println(err.Error())
			return errors.New("internal server error")
		}
		isPasswordValid := false
		passwordByte := []byte(*Password)
		isPasswordValid = passwords.PasswordMatch(*usersList.Password, passwordByte)
		if isPasswordValid == true {
			//succesful login
			return nil
		}
	}

	return errors.New("internal server error")
}
