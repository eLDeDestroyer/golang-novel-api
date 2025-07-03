package seeders

import (
	"e-novel/model"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {
	for i := 0; i < 6; i++ {

	password := "12345678"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	if err != nil {
		fmt.Println("Error hashing password:", err)
	}

		data := model.User{
			Username: "username" + strconv.Itoa(i),
			Name: "name" + strconv.Itoa(i),
			Password: string(hashedPassword),
		}


		db.Table("users").Create(&data)
	}

	return nil
}