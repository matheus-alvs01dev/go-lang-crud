package routes

import (
	"ApiGo/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.Index).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.Show).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.Store).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.Destroy).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.Edit).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
