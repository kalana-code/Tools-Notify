package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/notify/auth"
	"github.com/notify/genaral/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

)


// AddUser function save user in db
func AddUser(User model.UserInfo) model.InnerResponse {

	errIns := InsertUser(User)

	if errIns!=nil {
		log.Println("ERROR: DataBase Insert Query Excuting Process Failed. ")
		if strings.Contains(errIns.Error(), "Email_UNIQUE") {
			return errorDuplicateEmail
		}
		return errorDataBase	
	}

	log.Println("INFO: User has been inserted successfully.")
	return stateCreated

}

// FindUser used for find a user
func FindUser(User model.UserLogin) model.InnerResponse {
	resp := model.InnerResponse{}
	userInfo, err := GetUserInfo(User.Email)
	if err != nil {
		log.Println("ERROR: Email address not found", err)
		resp.Status = http.StatusForbidden
		resp.Message = "Email address not found"
		return resp
	}
	expiresAt := time.Now().Add(time.Hour).Unix()

	errEncry := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(User.Password))
	if errEncry != nil && errEncry == bcrypt.ErrMismatchedHashAndPassword {
		log.Println("ERROR: Invalid login credentials. Please try again", errEncry)
		resp.Status = http.StatusForbidden
		resp.Message = "Invalid login credentials. Please try again"
		return resp
	}

	JwtToken := &auth.Token{
		UserId : userInfo.UserId,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Email:     userInfo.Email,
		Role:      userInfo.Role,

		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), JwtToken)
	// sign token
	tokenString, error := token.SignedString([]byte(os.Getenv("jwtSecret")))
	
	// token generated time
	Issue_time := time.Now() 

	if error != nil {
		log.Println("ERROR: Token signing process failed", error)
		resp.Status = http.StatusInternalServerError
		resp.Message = "token signing process failed"
		return resp
	}
	resp.Status = http.StatusOK
	resp.Message = "logged in successfully"
	resp.Data = map[string]interface{}{"token": tokenString, "IssueAt": Issue_time}
	log.Println("INFO: Logged in successfully")

	return resp
}




