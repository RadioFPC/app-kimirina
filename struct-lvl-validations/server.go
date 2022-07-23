package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

// User contains user information.
type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `binding:"required,email"`
}

// UserStructLevelValidation contains custom struct level validations that don't always
// make sense at the field validation level. For example, this function validates that either
// FirstName or LastName exist; could have done that with a custom field validation but then
// would have had to add it to both fields duplicating the logic + overhead, this way it