package handlers

import (
	"log"
	"net/http"

	"github.com/bugjoe/family-tree-backend/persistence"

	"github.com/bugjoe/family-tree-backend/utils"

	"github.com/bugjoe/family-tree-backend/models"
)

// GetPersons returns all persons
func GetPersons(response http.ResponseWriter, request *http.Request) {

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

	_, err = persistence.UpsertPerson(person)
	if err != nil {
		log.Println("ERROR: Can't upsert person:", err)
		http.Error(response, "Error while upserting person", 500)
		return
	}

	response.WriteHeader(200)
}
