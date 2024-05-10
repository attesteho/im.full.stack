package main

import (
	"fmt"
	"net/http"
	"tp1/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialisation du routeur Gorilla Mux
	router := mux.NewRouter()

	// Définition des routes
	router.HandleFunc("/", controllers.HomeHandler)
	router.HandleFunc("/services", controllers.ServicesHandler)
	router.HandleFunc("/projects", controllers.ProjectsHandler)
	router.HandleFunc("/contact", controllers.ContactHandler)

	// Dossier statique pour les fichiers CSS et images
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Démarrage du serveur web
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
