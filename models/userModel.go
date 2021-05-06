package models

type UserModel struct {
	Id       int    `json: "id" validate: "required"`
	Username string `json: "username" validate: "required"`
	Password string `json: "password" validate: "required"`
	Name     string `json: "name" validate: "required"`
}
