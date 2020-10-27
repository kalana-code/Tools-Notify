package model

import (
	"net/http"
	"os"
)

// User is model is used for register
type User struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	BirthDay  string `json:"BirthDay"`
}

//InnerResponse used for exchange info between utils and service file
type InnerResponse struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

//Response is used for arrange response in standed way
type Response struct {
	Application string
	Version string
	Status  string
	Code    int
	Message string
	Data    map[string]interface{}
}

// Default set default value to the response
func (obj *Response) Default() {
	obj.Application = os.Getenv("application_name")
	obj.Version = os.Getenv("application_version")
}

// BadRequest set as Bad Request
func (obj *Response) BadRequest() {
	obj.Code = http.StatusBadRequest
	obj.Status = "Failed"
	obj.Message = "Bad Request"
}

// InternalServerError set Internal server error  Request
func (obj *Response) InternalServerError() {
	obj.Code = http.StatusInternalServerError
	obj.Status = "Failed"
	obj.Message = "Internal Server Error"
}

func (obj *Response) RestrictedError() {
	obj.Code = http.StatusForbidden
	obj.Status = "Failed"
	obj.Message = "Restricted "
}
// UserLogin is model is used for login process
type UserLogin struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// UserInfo is used for fetch user data from database
type UserInfo struct {
	UserId    int    `json:"UserId"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Role      string `json:"Role"`
	IsActive  bool 	 `json:"IsActive"`
}
