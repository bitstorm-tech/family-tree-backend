package persistence

import (
	"fmt"
	"log"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type bindingVariables map[string]interface{}

var (
	arangoClient driver.Client
	arangoDatabase driver.Database
	arangoCollections = make(map[string]driver.Collection)
	arangoGraph driver.Graph
	arangoEdgeCollections = make(map[string]driver.Collection)
	databaseName = "family-tree"
)

func getArangoClient() (driver.Client, error) {
	if arangoClient == nil {
		log.Println("Initialize ArangoDB client")
		endpoint := "http://localhost:8529"

		connectionConfig := http.ConnectionConfig{
			Endpoints: []string{endpoint},
		}

		connection, err := http.NewConnection(connectionConfig)
		if err != nil {
			return nil, err
		}

		clientConfig := driver.ClientConfig{
			Connection:     connection,
			Authentication: driver.BasicAuthentication("root", "root"),
		}

		arangoClient, err = driver.NewClient(clientConfig)
		if err != nil {
			return nil, err
		}
	}

	return arangoClient, nil
}

func getArangoDatabase() (driver.Database, error) {
	if arangoDatabase == nil {
		log.Println("Initialize ArangoDB database")
		client, err := getArangoClient()
		if err != nil {
			return nil, err
		}

		database, err := client.Database(nil, databaseName)
		if err != nil {
			return nil, err
		}

		arangoDatabase = database
	}

	return arangoDatabase, nil
}

func getArangoCollection(name string) (driver.Collection, error) {
	if arangoCollections[name] == nil {
		log.Println("Initialize ArangoDB collection:", name)
		database, err := getArangoDatabase()
		if err != nil {
			return nil, err
		}

		collection, err := database.Collection(nil, name)
		if err != nil {
			return nil, err
		}
		arangoCollections[name] = collection
	}

	return arangoCollections[name], nil
}

func getArangoGraph() (driver.Graph, error) {
	if arangoGraph == nil {
		log.Println("Initialize ArangoDB graph")
		database, err := getArangoDatabase()
		if err != nil {
			return nil, err
		}

		graph, err := database.Graph(nil, databaseName)
		if err != nil {
			return nil, err
		}

		arangoGraph = graph
	}

	return arangoGraph, nil
}

func getArangoEdgeCollection(name string) (driver.Collection, error) {
	if arangoEdgeCollections[name] == nil {
		log.Println("Initialize ArangoDB edge collection:", name)
		graph, err := getArangoGraph()
		if err != nil {
			return nil, err
		}

		edgeCollection, _, err := graph.EdgeCollection(nil, name)
		arangoEdgeCollections[name] = edgeCollection
	}

	return arangoEdgeCollections[name], nil
}

func createEdge(from driver.DocumentID, to driver.DocumentID, collection string) error {
	database, err := getArangoDatabase()
	if err != nil {
		return err
	}

	bindingVariables := bindingVariables{
		// "collection": []byte(collection),
		"from": from.String(),
		"to":   to.String(),
	}

	query := fmt.Sprintf("FOR edge IN %s FILTER edge._from == @from && edge._to == @to RETURN edge", collection)

	cursor, err := database.Query(nil, query, bindingVariables)
	if err != nil {
		return err
	}

	if cursor.HasMore() {
		return nil
	}

	query = fmt.Sprintf("INSERT { _from: @from, _to: @to } IN %s", collection)
	_, err = database.Query(nil, query, bindingVariables)
	if err != nil {
		return err
	}

	return nil
}
