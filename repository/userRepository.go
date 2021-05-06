package repository

import (
	"be-sharing-vision/driver"
	"be-sharing-vision/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type ResponseModel struct {
	Code    int    `json: "code" validate: "required"`
	Message string `json: "message" validate: "required"`
}

// HASH PASSWORD FUNC
func HashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashPass), err
}

// CREATE
func CreateNewUser(U *models.UserModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}

	db, err := driver.ConnectToDB()
	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	// LENGTH VALIDATION
	usernameLength := len(U.Username)
	passwordLength := len(U.Password)
	nameLength := len(U.Name)

	if usernameLength < 3 {
		Res := &ResponseModel{99, "Username too short!"}
		return Res
	}

	if passwordLength < 7 {
		Res := &ResponseModel{99, "Password too short!"}
		return Res
	}

	if nameLength < 3 {
		Res := &ResponseModel{99, "Name too short!"}
		return Res
	}

	hashedPassword, _ := HashPassword(U.Password)
	_, err = db.Exec(`INSERT INTO users (username, password, name) VALUES (?, ?, ?)`, U.Username, hashedPassword, U.Name)

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Create New User Failed"}
		return Res
	}

	defer db.Close()

	fmt.Println("Create Data Success!")
	Res = &ResponseModel{200, "Create New User Success"}
	return Res
}

// READ ALL
func ReadAllUsers(limit int, offset int) []models.UserModel {
	db, err := driver.ConnectToDB()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var result []models.UserModel
	items, err := db.Query(`SELECT id, username, password, name FROM users LIMIT ? OFFSET ?`, limit, offset)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.UserModel{}
		var err = items.Scan(&each.Id, &each.Username, &each.Password, &each.Name)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)
	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	return result
}

// READ BY ID
func ReadUserById(id int) []models.UserModel {
	db, err := driver.ConnectToDB()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var result []models.UserModel

	items, err := db.Query(`SELECT * FROM users WHERE id = ?`, id)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = models.UserModel{}
		var err = items.Scan(&each.Id, &each.Username, &each.Password, &each.Name)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}

	defer db.Close()

	return result
}

// UPDATE
func UpdateUser(U *models.UserModel, id int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}

	db, err := driver.ConnectToDB()
	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	hashedPassword, _ := HashPassword(U.Password)
	_, err = db.Exec(`UPDATE users SET username = ?, password = ?, name = ? WHERE id = ?`, U.Username, hashedPassword, U.Name, id)

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Update User Data Failed"}
		return Res
	}

	defer db.Close()

	fmt.Println("Update Data Success!")
	Res = &ResponseModel{200, "Update User Data Success"}
	return Res
}

// DELETE
func DeleteUser(id int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}

	db, err := driver.ConnectToDB()
	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	_, err = db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		fmt.Print(err.Error())
		Res = &ResponseModel{400, "Delete User Data Failed"}
		return Res
	}

	defer db.Close()

	fmt.Println("Delete Data Success!")
	Res = &ResponseModel{200, "Delete User Data Success"}
	return Res
}
