package auth

import "github.com/dgrijalva/jwt-go"

// Token is used handle jwt token
type Token struct {
	UserId int
	FirstName string
	LastName  string
	Gender    string
	Email     string
	Role      string
	*jwt.StandardClaims
}


// ApplicationKey used for tack application 
type ApplicationKey struct {
	ApplicationId string
	ApplicationName string
	*jwt.StandardClaims
}