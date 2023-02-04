package personality

import (
	"errors"

	"github.com/LucasRMP/golang-studies-rest-api/database"
)

func FindAll() []Personality {
	var personalities []Personality
	database.DB.Find(&personalities)
	return personalities
}

func FindById(id int) (Personality, error) {
	var personality Personality
	database.DB.First(&personality, id)
	if personality.Id != 0 {
		return personality, nil
	}

	return Personality{}, errors.New("Personality not found")
}

func Create(input CreatePersonalityDTO) (Personality, error) {
	personality := Personality{
		Name:      input.Name,
		Biography: input.Biography,
	}

	result := database.DB.Create(&personality)
	if result.Error != nil {
		return Personality{}, result.Error
	}

	return personality, nil
}

func Delete(id int) error {
	var personality Personality
	database.DB.First(&personality, id)

	if personality.Id == 0 {
		return errors.New("Personality not found")
	}

	database.DB.Delete(&personality)
	return nil
}

func Update(id int, input UpdatePersonalityDTO) (Personality, error) {
	var personality Personality
	database.DB.First(&personality, id)

	if personality.Id == 0 {
		return Personality{}, errors.New("Personality not found")
	}

	if input.Name != "" {
		personality.Name = input.Name
	}
	if input.Biography != "" {
		personality.Biography = input.Biography
	}

	database.DB.Save(&personality)

	return personality, nil
}
