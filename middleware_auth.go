package main

import (
	"fmt"
	"net/http"

	"github.com/RishabhSharma17/rssaggregator/internal/database"
	"github.com/RishabhSharma17/rssaggregator/internal/database/auth"
)

type AuthHandler func(http.ResponseWriter , *http.Request , database.User) 

func (config *apiConfig) middlewareAuth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey,err := auth.GetAPIKey(r.Header)
		
		if err!=nil{
			respondWithError(w,403,fmt.Sprintf("Auth error : %v ",err))
			return
		}

		user,err := config.DB.GetUserByAPIKey(r.Context(),apiKey)
		if err!= nil {
			respondWithError(w,400,fmt.Sprintf("couldn't get user:%s",err))
			return
		}

		handler(w,r,user)
	}
}