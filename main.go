package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"github.com/realr3fo/tkai_circles_ball/app"
	"github.com/realr3fo/tkai_circles_ball/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	//router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	//router.HandleFunc("/api/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts
	router.HandleFunc("/api/circle/area", controllers.GetCircleArea).Methods("POST")
	router.HandleFunc("/api/tube/volume", controllers.GetTubeVolume).Methods("POST")
	router.HandleFunc("/api/ball/volume", controllers.GetBallVolume).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8002" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
