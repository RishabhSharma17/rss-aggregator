package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RishabhSharma17/rssaggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func main(){
	godotenv.Load()

	PortNumber := os.Getenv("PORT")
	if PortNumber == "" {
		//It is used to exit the program with the execution and return some error
		log.Fatal("PORT is not provided in the environment variable")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == ""{
		log.Fatal("PORT is not provided in the environment variable")
	}

	conn,err := sql.Open("postgres",dbUrl)
	if err!=nil{
		log.Fatal("can't connect to the database:",err)
	}


	apicfg := apiConfig{
		DB: database.New(conn),
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
	routerV1.Post("/user",apicfg.handleCreateUser)
	routerV1.Get("/user",apicfg.middlewareAuth(apicfg.handleGetUser))
	routerV1.Post("/feed",apicfg.middlewareAuth(apicfg.handleCreateFeed))
	routerV1.Get("/feed",apicfg.handleGetFeed)
	routerV1.Post("/feed_follow",apicfg.middlewareAuth(apicfg.handleCreateFeedFollows))
	routerV1.Get("/feed_follow",apicfg.middlewareAuth(apicfg.handleGetFeedFollows))
	routerV1.Delete("/feed_follow/{feedfollowID}",apicfg.middlewareAuth(apicfg.handleDeleteFeedFollow))

	router.Mount("/v1", routerV1)

	srvr := &http.Server{
		Handler:router,
		Addr:":"+PortNumber,
	}

	fmt.Printf("Server is running on port %s", PortNumber)
	err = srvr.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

}