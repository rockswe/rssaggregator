package main

import (
	"encoding/json"
	"fmt"
	"net/http" a

	"github.com/google/uuid"
	"github.com/rockswe/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Printf("Error parsing JSON:", err))
		return
	}

	apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New().String(),
	})

	respondwithJSON(w, 200, struct{}{})
}
