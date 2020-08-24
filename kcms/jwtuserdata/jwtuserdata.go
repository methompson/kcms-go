package jwtuserdata

import "github.com/dgrijalva/jwt-go"

// JWTUserData provides a unified data type for JWT information (both for encoding and decoding)
// This struct is in its own package to prevent import cycles
type JWTUserData struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	UserType  string `json:"userType"`
	UserMeta  string `json:"userMeta"`
	jwt.StandardClaims
}
