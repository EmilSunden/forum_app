package controllers

import (
	"fmt"
	"net/http"
)

func LoginController(res http.ResponseWriter, r *http.Request) {
	// LoginController is the function that handles the login logic for the application
	// Login logic here
	fmt.Println("LoginController")
	res.Header().Set("Content-Type", "application/json")

	res.Write([]byte("LoginController"))
}
