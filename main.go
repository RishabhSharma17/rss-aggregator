package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	PortNumber := os.Getenv("PORT")

	if PortNumber == "" {
		//It is used to exit the program with the execution and return some error
		log.Fatal("PORT is not provided in the environment variable")
	}

	router := chi.NewRouter()
	routerV1 := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routerV1.Get("/healthz",handleReadiness)
	routerV1.Get("/err",handleError)

	router.Mount("/v1", routerV1)

	srvr := &http.Server{
		Handler:router,
		Addr:":"+PortNumber,
	}

	fmt.Printf("Server is running on port %s", PortNumber)
	err := srvr.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

}