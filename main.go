package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest-go-demo/controllers"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}

func initDB() *gorm.DB {
	// db, err := gorm.Open("mysql", "docker:password@tcp(godockerDB)/godocker")

	user := os.Getenv("admin")
	pass := os.Getenv("Control123")
	host := os.Getenv("127.0.0.1")
	dbname := os.Getenv("learning_demo")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)

	db, err := gorm.Open("mysql", connection)

	// config :=
	// 	database.Config{
	// 		ServerName: "localhost:3306",
	// 		User:       "admin",
	// 		Password:   "Control@123",
	// 		DB:         "learning_demo",
	// 	}

	// connectionString := database.GetConnectionString(config)
	// err := database.Connect(connectionString)
	if err != nil {
		fmt.Println("could'n connect!")
		panic(err.Error())
	}
	// database.Migrate(&entity.Person{})
	return db
}
