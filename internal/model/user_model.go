package model

import "time"

type User struct {
	ID           int       `bson:"_id"`
	First_name   string    `bson:"first_name"`
	Last_name    string    `bson:"last_name"`
	Username     string    `bson:"username" unique:"true"`
	Password     string    `bson:"password"`
	Phone_number string    `bson:"phone_number" unique:"true"`
	Email        string    `bson:"email" unique:"true"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}

func NewUserModel(firstName, lastName, username, password, phoneNumber, email string) User {
	return User{
		First_name:   firstName,
		Last_name:    lastName,
		Username:     username,
		Password:     password,
		Phone_number: phoneNumber,
		Email:        email,
	}
}
