/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

type CatForm struct {
	FormBase
	Name string `json:"name" validate:"required,min=2,max=50"`
	Type string `json:"type" validate:"required,min=2,max=80"`
}

func (form *CatForm) MapToModel() *Cat {
	return &Cat{
		Name: form.Name,
		Type: form.Type,
	}
}
