package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RishabhSharma17/rssaggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handleCreateFeed(w http.ResponseWriter , r *http.Request,user database.User){
	type parameter struct{
		Name string `json:"name"`
		Url string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameter{}

	err := decoder.Decode(&params)

	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Error while parsing json:%v",err))
		return
	}

	feed,err := config.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	})
	
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("couldn't create feed:%s",err))
        return
	}

	respondWithJson(w,201,DatabaseFeedToFeed(feed))
}

func (config *apiConfig) handleGetFeed(w http.ResponseWriter , r *http.Request){
	feed,err := config.DB.GetFeeds(r.Context())

	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("couldn't get feed:%s",err))
        return
	}


	respondWithJson(w,201,DatabaseFeedsToFeeds(feed))
}