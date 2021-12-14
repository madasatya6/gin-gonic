// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

//Login credential
type LoginCredentials struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=5"`
}
