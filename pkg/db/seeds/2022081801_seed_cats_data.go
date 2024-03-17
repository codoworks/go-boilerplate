/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package seeds

import (
	"github.com/codoworks/go-boilerplate/pkg/db/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {

	var s = &gormigrate.Migration{}
	s.ID = "2022081801_seed_cats_data"

	s.Migrate = func(db *gorm.DB) error {

		var err error
		var cats []*models.Cat = []*models.Cat{}

		cats = append(cats, &models.Cat{
			Name: "Kitty",
			Type: "Persian",
		})

		cats = append(cats, &models.Cat{
			Name: "Whiskers",
			Type: "Siamese",
		})

		cats = append(cats, &models.Cat{
			Name: "Fluffy",
			Type: "Ragdoll",
		})

		cats = append(cats, &models.Cat{
			Name: "Tom",
			Type: "Siamese",
		})

		cats = append(cats, &models.Cat{
			Name: "Jerry",
			Type: "Siamese",
		})

		cats = append(cats, &models.Cat{
			Name: "Garfield",
			Type: "Ragdoll",
		})

		cats = append(cats, &models.Cat{
			Name: "Sylvester",
			Type: "Persian",
		})

		cats = append(cats, &models.Cat{
			Name: "Felix",
			Type: "Birman",
		})

		cats = append(cats, &models.Cat{
			Name: "Heathcliff",
			Type: "Common Cat",
		})

		cats = append(cats, &models.Cat{
			Name: "Snowball",
			Type: "Ragdoll",
		})

		cats = append(cats, &models.Cat{
			Name: "Salem",
			Type: "Birman",
		})

		cats = append(cats, &models.Cat{
			Name: "Cheshire",
			Type: "Ragdoll",
		})

		for _, cat := range cats {
			err = cat.Save()
			if err != nil {
				logFail(s.ID, err)
				return err
			}
		}

		logSuccess(s.ID)
		return nil
	}

	AddSeed(s)
}
