package router

import (
	"TP-API-Spotify/controller"
	"net/http"
)

// New crée et retourne un nouvel objet ServeMux configuré avec les routes de l'application
func New() *http.ServeMux {
	mux := http.NewServeMux()

	// Routes de ton app
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/Damso", controller.Damso)
	mux.HandleFunc("/Laylow", controller.Laylow)
	mux.HandleFunc("/contact", controller.Contact)

	// Ajout des fichiers statiques
	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	return mux
}
