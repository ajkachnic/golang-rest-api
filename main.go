package main

import (
	"fmt"
	"log"
	"net/http"
	//"encoding/json"
	"github.com/gorilla/mux"
	//"io/ioutil"
	"github.com/ajkachnic/golang-rest-api/routes"
)
type Article struct {
	Id		string 	`json:"Id"`
	Title 	string 	`json:"Title"`
	Desc 	string 	`json:"desc"`
	Content string 	`json:"content"`
}

var Articles []Article

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// Route Handling
	myRouter.HandleFunc("/", routes.HomePage)
	myRouter.HandleFunc("/articles", routes.ReturnAllArticles)

	myRouter.HandleFunc("/article", routes.CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}&{title}&{desc}&{content}", routes.CreateNewArticleArgBased).Methods("POST")
	myRouter.HandleFunc("/article/{id}/{title}/{desc}/{content}", routes.CreateNewArticleArgBased).Methods("POST")

	myRouter.HandleFunc("/article/{id}", routes.DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", routes.UpdateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", routes.ReturnSingleArticle)

	log.Fatal(http.ListenAndServe(":3030", myRouter))
}

func main() {
	routes.Main()
	fmt.Println("Running API Server")
	handleRequests()
}
