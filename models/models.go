package models

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
