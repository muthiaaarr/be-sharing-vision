package service

import (
	"be-sharing-vision/models"
	"be-sharing-vision/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ResponseModel struct {
	Code    int    `json: "code" validate: "required"`
	Message string `json: "message" validate: "required"`
}

// CREATE
func CreateNewUser(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.UserModel)
	if err := c.Bind(U); err != nil {
		return nil
	}

	Res = (*ResponseModel)(repository.CreateNewUser(U))

	return c.JSON(http.StatusOK, Res)
}

// READ ALL
func ReadAllUsers(c echo.Context) error {
	limit := c.Param("limit")
	limitData, _ := strconv.Atoi(limit)
	offset := c.Param("offset")
	offsetData, _ := strconv.Atoi(offset)
	result := repository.ReadAllUsers(limitData, offsetData)

	return c.JSON(http.StatusOK, result)
}

// READ BY ID
func ReadUserById(c echo.Context) error {
	id := c.Param("id")
	data, _ := strconv.Atoi(id)
	result := repository.ReadUserById(data)

	return c.JSON(http.StatusOK, result)
}

// UPDATE
func UpdateUser(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.Param("id")
	data, _ := strconv.Atoi(id)
	U := new(models.UserModel)

	if err := c.Bind(U); err != nil {
		return nil
	}

	Res = (*ResponseModel)(repository.UpdateUser(U, data))

	return c.JSON(http.StatusOK, Res)
}

// DELETE
func DeleteUser(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.Param("id")
	data, _ := strconv.Atoi(id)
	Res = (*ResponseModel)(repository.DeleteUser(data))

	return c.JSON(http.StatusOK, Res)
}
