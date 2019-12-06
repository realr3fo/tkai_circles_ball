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
