/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package helpers

import "github.com/go-playground/validator/v10"

type GoogleValidator struct {
	validator *validator.Validate
}

func (gv *GoogleValidator) Validate(s interface{}) error {
	return gv.validator.Struct(s)
}
