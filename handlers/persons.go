package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bugjoe/family-tree-backend/persistence"

	"github.com/bugjoe/family-tree-backend/utils"

	"github.com/bugjoe/family-tree-backend/models"
)

// GetPersons returns all persons
func GetPersons(response http.ResponseWriter, request *http.Request) {
	persons, err := persistence.GetAllPersons()
	if err != nil {
		log.Println("ERROR: Can't get all persons:", err)
		http.Error(response, "Error while getting all persons", 500)
		return
	}

	json, err := json.Marshal(persons)
	if err != nil {
		log.Println("ERROR: Can't convert persons to JSON []byte:", err)
		http.Error(response, "Error while converting persons into JSON", 500)
		return
	}

	response.Write(json)
}

// UpsertPerson inserts a new person if it does not exist in the database
// yet or updates an already existing person.
func UpsertPerson(response http.ResponseWriter, request *http.Request) {
	person := new(models.Person)
	err := utils.ExtractFromRequest(request, person)
	if err != nil {
		log.Println("ERROR: Can't extract person from request:", err)
		http.Error(response, "Error while parsing request body", 500)
		return
	}

	person, err = persistence.UpsertPerson(person)
	if err != nil {
		log.Println("ERROR: Can't upsert person:", err)
		http.Error(response, "Error while upserting person", 500)
		return
	}

	json, err := json.Marshal(person)
	if err != nil {
		log.Println("ERROR: Can't convert person to JSON []byte", err)
		http.Error(response, "Error while converting person into JSON", 500)
		return
	}

	response.Write(json)
}
