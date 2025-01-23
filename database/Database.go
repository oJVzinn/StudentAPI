package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"studentapi/person"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(person.Person{})
	if err != nil {
		log.Fatal(err)
	}
}

func Create(person *person.Person) error {
	result := db.Create(person)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Remove(person *person.Person) error {
	result := db.Delete(person)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAll() (*[]person.Person, error) {
	var persons []person.Person
	result := db.Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}

	return &persons, nil
}

func Get(id int) (*person.Person, error) {
	var person person.Person
	result := db.First(&person, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &person, nil
}

func Update(person *person.Person) error {
	result := db.Save(person)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
