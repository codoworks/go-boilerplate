/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package context

type ApiResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Payload interface{} `json:"payload"`
}

type FieldValidationError struct {
	Namespace string      `json:"namespace"`
	Field     string      `json:"field"`
	Error     string      `json:"error"`
	Kind      string      `json:"kind"`
	Value     interface{} `json:"value"`
}
