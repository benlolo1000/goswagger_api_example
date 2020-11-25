package registerUser

import (
	"errors"
	"log"
	"medicarePartD/bus/db/userDatabase"
	"medicarePartD/bus/passwords"
	"regexp"
)

func RegisterUser(Email *string, Password *string, ConfirmPassword *string) error {

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

	//check for pre-existing account with provided email
	rows, err := userDb.Query("SELECT COUNT(*) FROM Users WHERE Email=?", *Email)
	if err != nil {
		log.Println(err.Error())
		return errors.New("internal server error")
	}
	for rows.Next() {
		var count int
		err := rows.Scan(&count)
		if err != nil {
			return errors.New("internal server error")
		}
		if count != 0 {
			return errors.New("conflict")
		}

	}
	//validate password
	if *Password != *ConfirmPassword {
		return errors.New("not acceptable")
	} else {
		passLength := len(*Password)
		if passLength <= 8 && passLength >= 32 && *ConfirmPassword != *Password {
			return errors.New("not acceptable")
		} else {
			// insert into database
			passwordByte := []byte(*Password)
			hashedPassword := passwords.HashAndSalt(passwordByte)
			_, err := userDb.Exec("INSERT INTO Users (Email, Pass) VALUES (?, ?)", *Email, hashedPassword)
			if err != nil {
				log.Println(err.Error())
				return errors.New("internal server error")
			}
		}
	}
	//succesful registration
	return nil
}
