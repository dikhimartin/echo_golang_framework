package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"../models"
	"github.com/labstack/echo"
)

var users = models.Users{
	models.User{
		ID:           1,
		NamaDepan:    "Jojo",
		NamaBelakang: "Si",
		Email:        "j@gmail.com",
	},
	{
		ID:           1,
		NamaDepan:    "Jojo",
		NamaBelakang: "Si",
		Email:        "j@gmail.com",
	},
}

//Adduser
func AddUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	newUser := models.User{
		ID:           id,
		NamaDepan:    c.FormValue("namaDepan"),
		NamaBelakang: c.FormValue("namaBelakang"),
		Email:        c.FormValue("email"),
	}

	users = append(users, newUser)

	fmt.Printf("%v", users)

	return c.JSON(http.StatusCreated, newUser)
}

//Getusers
func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
