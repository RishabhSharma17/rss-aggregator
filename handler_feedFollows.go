package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RishabhSharma17/rssaggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (config *apiConfig) handleCreateFeedFollows(w http.ResponseWriter , r *http.Request, user database.User){
	type parameter struct{
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameter{}

	err := decoder.Decode(&params)

	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Error while parsing json:%v",err))
		return
	}

	feedfollow,err := config.DB.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		FeedID: params.FeedID,
		UserID: user.ID,
		CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
		ID: uuid.New(),
	})
	
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("couldn't create feed:%s",err))
        return
	}

	respondWithJson(w,201,DatabaseFeedFollowToFeedFollow(feedfollow))
}

func (config *apiConfig) handleGetFeedFollows(w http.ResponseWriter , r *http.Request , user database.User) {
	feedfollow,err := config.DB.GetFeedFollow(r.Context(),user.ID)
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("couldn't get feed follows:%s",err))
        return
	}
	respondWithJson(w,201,DatabaseFeedFollowsToFeedFollows(feedfollow))
}

func (config *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter,r *http.Request, user database.User){
	feedfollowstr := chi.URLParam(r, "feedfollowID")

	feedfollowid , err := uuid.Parse(feedfollowstr)
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("couldn't parse feedfollowid:%v",err))
		return
	}

	err = config.DB.DeleteFeedFollow(r.Context(),database.DeleteFeedFollowParams{
		ID: feedfollowid,
        UserID: user.ID,
	})

	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("couldn't delete feedfollow:%s",err))
        return
	}
	respondWithJson(w,200,struct{}{})
}