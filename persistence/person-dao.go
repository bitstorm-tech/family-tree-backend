package persistence

import "github.com/bugjoe/family-tree-backend/models"

// UpsertPerson either inserts the person if it does not already
// exist in the database or updates an existing person.
func UpsertPerson(person *models.Person) (*models.Person, error) {
	collection, err := getArangoCollection("persons")
	if err != nil {
		return nil, err
	}

	if len(person.Key) > 0 {
		_, err := collection.UpdateDocument(nil, person.Key, person.Payload)
		if err != nil {
			return nil, err
		}
	} else {
		meta, err := collection.CreateDocument(nil, person.Payload)
		if err != nil {
			return nil, err
		}

		person.Key = meta.Key
	}

	return person, nil
}

// GetAllPersons returns all persons from the database.
func GetAllPersons() ([]models.Person, error) {
	database, err := getArangoDatabase()
	if err != nil {
		return nil, err
	}

	query := "FOR p IN persons RETURN p"
	cursor, err := database.Query(nil, query, nil)
	if err != nil {
		return nil, err
	}

	persons := []models.Person{}
	for cursor.HasMore() {
		person := new(models.Person)
		_, err := cursor.ReadDocument(nil, &person.Payload)
		if err != nil {
			return nil, err
		}

		persons = append(persons, *person)
	}

	return persons, nil
}
