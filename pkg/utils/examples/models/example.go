/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

// README:-
//
// Defining a new model should be as simple as duplicating this file
// and renaming the struct.
// The file name to match the new model name.
// Simply find and replace the word "example" with the new model name.
// make sure to maintain the case sensitivity of the names "example"
// and "Example".
//
// Once that is done, the new model should be ready to use.
// you can now add new fields to the model and any additional model
// specific methods that you may need.
//

// The first thing you need is a pointer to the struct
// This is needed to enable the handler from accessing
// the struct methods
//
// when creating a new model, make sure to maintain this pattern
/*
var example *Example = &Example{}
*/

// ExampleModel returns a pointer to the Example
// struct. This is the only way to access the
// Example struct to ensure a singleton pattern
// is enforced.
//
// This can be used to access the Example struct
// in the following way, which is quite important
// for the handers to avoid creating new instances
// of the struct every time only to call a method:
//
// example, err := ExampleModel().Find(id)
// examples, error := ExampleModel().FindAll()
// err := ExampleModel().Save()
// err := ExampleModel().Update()
// err := ExampleModel().Delete(id)
// ... and so on
/*
func ExampleModel() *Example {
	return example
}
*/

// This is the database model.
// it corresponds to the database table
// named "examples".
// For references on how to define the struct,
// see the gorm documentation at gorm.io

// The struct should be defined as follows:
/*
type Example struct {
	ModelBase
	Title       string `gorm:"size:255"`
	Description string `gorm:"size:255"`
	Views       int    `gorm:""`
	Slug        string `gorm:"size:255"`
	Content     string `gorm:"size:255"`

	// you can also add additional fields here
}
*/

// The struct should have the following method:
// MapToForm() *ExampleForm
//
// It returns a response form that can be mapped
// to json and returned to the client
//
// The purpose of this method is to allow for
// multiple forms to be returned, that way you
// can reuse the model to generate sanitized
// responses for different use cases.
/*
func (model *Example) MapToForm() *ExampleForm {
	form := &ExampleForm{
		Title:       model.Title,
		Description: model.Description,
		Views:       model.Views,
		Slug:        model.Slug,
		Content:     model.Content,

		// you can also add additional fields here
	}
	form.ID = model.ID
	form.CreatedAt = model.CreatedAt
	form.UpdatedAt = model.UpdatedAt
	return form
}
*/

// The struct should have the following methods:
// FindAll() ([]*Example, error)
//
// This will return a list of all the records in
// the "examples" table in the database.
/*
func (model *Example) FindAll() (models []*Example, err error) {
	result := db.Model(model).Find(&models)
	return models, result.Error
}
*/

// The struct should have the following methods:
// FindMany(ids []string) ([]*Example, error)
//
// This will return a list of all the records in
// the "examples" table in the database that match
// the ids provided in the ids slice.
/*
func (model *Example) FindMany(ids []string) (models []*Example, err error) {
	result := db.Model(model).Find(&models, ids)
	return models, result.Error
}
*/

// The struct should have the following methods:
// Find(id string) (*Example, error)
//
// This will return a single record from the "examples"
// table in the database that matches the id provided.
/*
func (model *Example) Find(id string) (m *Example, err error) {
	result := db.Model(model).Where("ID=?", id).First(&m)
	return m, result.Error
}
*/

// The struct should have the following methods:
// Save() error
//
// This will save the current instance of the struct
// to the "examples" table in the database.
/*
func (model *Example) Save() error {
	return db.Model(model).Create(&model).Error
}
*/

// The struct should have the following methods:
// Update() error
//
// This will update the current instance of the struct
// in the "examples" table in the database.
/*
func (model *Example) Update() error {
	return db.Model(model).Updates(&model).Error
}
*/

// The struct should have the following methods:
// Delete(id string) error
//
// This will delete the record from the "examples" table
// in the database that matches the id provided.
/*
func (model *Example) Delete(id string) error {
	return db.Model(model).Where("ID=?", id).Delete(&model).Error
}
*/
