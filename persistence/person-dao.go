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
