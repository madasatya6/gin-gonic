// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Setting Model -.
type Setting struct {
	ID			int	   `form:"id" 	  		query:"id" 			json:"id" 			binding:"required"`
	Name      	string `form:"name"   		query:"name" 		json:"name" 		binding:"required"`
	Value 		string `form:"value"  		query:"value" 		json:"value"		binding:"required"`
	FieldType   string `form:"field_type" 	query:"field_type" 	json:"field_type" 	binding:"required"`
	Tab 		string `form:"tab" 			query:"tab" 		json:"tab" 			binding:"required"`
}
