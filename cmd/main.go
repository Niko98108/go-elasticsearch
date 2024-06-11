package main

import (
	"encoding/json"
	con "go-elasticsearch/connection"
	"go-elasticsearch/usercase"
	"log"
)

func main() {

	con.Connector()
	cl := usercase.Common{}

	// Instantiate new Elasticsearch documents from the ElasticIndex struct
	newDocs := []usercase.ElasticIndex{
		{FieldStr: "hello hello", FieldInt: 42, FieldBool: true},
		{FieldStr: "hello world", FieldInt: 777, FieldBool: false},
		{FieldStr: "bonjour monde", FieldInt: 1234, FieldBool: true},
	}
	byteDoc, err := json.Marshal(newDocs)
	if err != nil {
		log.Fatalf("Error marshaling document: %s", err)
	}

	cl.IndexDocument("example", "1111", string(byteDoc))
}
