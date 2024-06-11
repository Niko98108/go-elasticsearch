package usercase

import "github.com/elastic/go-elasticsearch/v8"

type Common struct {
	client *elasticsearch.Client
}

// Film represents a movie with some properties.
type ElasticIndex struct {
	FieldStr  string `json:"field str"`
	FieldInt  int    `json:"field int"`
	FieldBool bool   `json:"field bool"`
}
