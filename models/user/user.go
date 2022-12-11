package user

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `idx:"{email},unique" bson:"email,omitempty" json:"email" binding:"required,email"`
	Username string             `idx:"unique" bson:"username,omitempty" binding:"required,min=5"`
	Password string             `bson:"password,omitempty" binding:"required,containsany=!@#$%^&*,min=8"`
	Apps     int                `bson:"apps" json:"apps" binding:"required"`
}

// GetToken creates a JWT Token for recognizing the current user
// Currently uses the email of the user
// TODO: Add more fields such as Expiration Time,Username
func (user *User) GetToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": string(user.Email),
	})

	tokenString, err := token.SignedString([]byte("secretkey"))
	return tokenString, err
}
