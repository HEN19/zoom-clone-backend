package routes

import (
	"self-improvement/zoom-clone-backend/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/api/users/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/users/login", controllers.LoginUser).Methods("POST")

	// Room routes
	router.HandleFunc("/api/rooms/create", controllers.CreateRoom).Methods("POST")
	router.HandleFunc("/api/rooms/{id}", controllers.GetRoom).Methods("GET")
	router.HandleFunc("/api/signal", controllers.SignalingHandler).Methods("GET")


	return router
}
