package models

type User struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Role     string `json:"role" bson:"role"`
	Password string `json:"password" bson:"password"`
}

func NewUser(name string, email string, role string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Role:     role,
		Password: password,
	}
}
