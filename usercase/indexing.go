package usercase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
)


func (es *Common) IndexDocument(indexName string, documentID string, document interface{}) {
	jsonDoc, err := json.Marshal(document)
	if err != nil {
		log.Fatalf("Error marshaling document: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: documentID,
		Body:       strings.NewReader(string(jsonDoc)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error indexing document ID=%s: %s", documentID, res.String())
	}

	fmt.Printf("Document ID=%s indexed successfully\n", documentID)
}
