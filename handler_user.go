package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RishabhSharma17/rssaggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameter{}

	err := decoder.Decode(&params)

	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Error while parsing json:%v",err))
		return
	}

	user,err := config.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("couldn't create user:%s",err))
		return
	}
	respondWithJson(w,200,DatabaseUserToUser(user))
}

func (config *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request){
	
}