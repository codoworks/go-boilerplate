/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

// README:-
//
// Defining a new form should be as simple as duplicating this file
// and renaming the struct.
// The file name to match the new model name, folloed by "Form".
// Simply find and replace the word "example" with the new model name.
// make sure to maintain the case sensitivity of the names "example"
// and "Example".
//
// Once that is done, the new form should be ready to use.
// you can now add new fields to the model and any additional model
// specific methods that you may need.
//
// Forms are also responsible on defining 4 things:
// 1. The fields that are expected to be received from the client
// 2. The fields that are expected to be returned to the client
// 3. json mapping
// 4. validation rules
//
// The purpose of having forms is to allow for
// multiple forms to be returned, that way you
// can reuse the model to generate sanitized
// responses for different use cases.

// Forms are simple structs that embed the FormBase struct
// and define the fields that are expected to be received
// from the client. The fields are then mapped to the
// model and validated before being used to create or
// update the model.
//
// FormBase would automatically provide the following
// fields:
// 1. ID
// 2. CreatedAt
// 3. UpdatedAt
// 4. DeletedAt # optional field
/*
type ExampleForm struct {
	FormBase
	Title       string `json:"name" validate:"required,min=2,max=50"`
	Description string `json:"description" validate:"required,min=2,max=255"`
	Views       int    `json:"views"`
	Slug        string `json:"slug"`
	Content     string `json:"content" content:"required,min=2,max=255"`

	// you can also add additional fields here
}
*/

// this method allows you to create a database model
// from a form containing the user input.
/*
func (form *ExampleForm) MapToModel() *Example {
	return &Example{
		Title:       form.Title,
		Description: form.Description,
		Views:       form.Views,
		Slug:        form.Slug,

		// you can also add additional fields here
	}
}
*/
