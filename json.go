package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func respondWithError(w http.ResponseWriter,code int, msg string){
	if code>499{
		log.Println("Responding with 5XX error: ",msg)
	}
	type ErrorResponse struct{
		Err string `json:"error"`
	}
	respondWithJson(w,code,ErrorResponse{
		Err:msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data,err := json.Marshal(payload) // this is for the conversion of payload to json
	if err!=nil{
		log.Println("Failed to marshal json Response : ",payload)
		w.WriteHeader(500)
        return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}