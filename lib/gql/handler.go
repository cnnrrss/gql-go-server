package gql

import (
	"encoding/json"
	"net/http"

	"github.com/graph-gophers/graphql-go"
)

type handler struct {
	schema *graphql.Schema
}

func NewHandler(schema *graphql.Schema) *handler {
	return &handler{
		schema: schema,
	}
}

func (h *handler) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	var params struct {
		Query string `json:"query"`
		OperationName string `json:"operationName"`
		Variables map[string]interface{} `json:"variables"`
	}

	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	response := h.schema.Exec(req.Context(), params.Query, params.OperationName, params.Variables)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}