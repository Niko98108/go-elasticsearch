package connection

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func Connector() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200", // this value should come as config parameter
		},
		// If your Elasticsearch instance requires authentication, uncomment and set these:
		// Username: "your_username",
		// Password: "your_password",
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Test the client by getting the cluster health
	res, err := es.Cluster.Health()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	fmt.Println(res)
}
