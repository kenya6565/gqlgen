package graph

import (
	"encoding/json"
	"os"

	"github.com/kenya6565/graphql_go/graph/model"
)

// type Resolver struct {
// 	query string
//  }

type Resolver struct {
	query struct {
	  Todos []*model.Todo `json:"todos"`
	}
}
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver() *Resolver {
	filename := "Query.json"

	// data = []byte
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Failed to read: " + filename)
	}
	
	resolver := Resolver{}
	// json string -> Go structに変換
	json.Unmarshal(data, &resolver.query)
	return &resolver
}