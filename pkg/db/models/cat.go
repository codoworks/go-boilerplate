/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

var cat *Cat = &Cat{}

func CatModel() *Cat {
	return cat
}

type Cat struct {
	ModelBase
	Name string `gorm:"size:255"`
	Type string `gorm:"size:255"`
}

func (model *Cat) MapToForm() *CatForm {
	form := &CatForm{
		Name: model.Name,
		Type: model.Type,
	}
	form.ID = model.ID
	form.CreatedAt = model.CreatedAt
	form.UpdatedAt = model.UpdatedAt
	return form
}

func (model *Cat) FindAll() (models []*Cat, err error) {
	result := db.Model(model).Find(&models)
	return models, result.Error
}

func (model *Cat) FindMany(ids []string) (models []*Cat, err error) {
	result := db.Model(model).Find(&models, ids)
	return models, result.Error
}

func (model *Cat) Find(id string) (m *Cat, err error) {
	result := db.Model(model).Where("ID=?", id).First(&m)
	return m, result.Error
}

func (model *Cat) Save() error {
	return db.Model(model).Create(&model).Error
}

func (model *Cat) Update() error {
	return db.Model(model).Updates(&model).Error
}

func (model *Cat) Delete(id string) error {
	return db.Model(model).Where("ID=?", id).Delete(&model).Error
}
