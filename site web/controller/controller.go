package controller

import (
	"TP-API-Spotify/api"
	"TP-API-Spotify/structure"
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

type PageData struct {
	Title     string
	Message   string
	Artist    string
	AlbumData []structure.Items
	Track     string
}

var Token *string

func Home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue 🎉",
	}
	// Récupération du token pour toute la session de l'utilisateur
	T := api.GetToken()
	if T.Error != "" {
		fmt.Println("Erreur lors de la récupération du token : ", T.Error, " ", T.ErrorDescription)
	} else {
		Token = &T.AccessToken
		fmt.Println("Token récupéré : ", *Token)
	}

	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Damso(w http.ResponseWriter, r *http.Request) {

	A := api.GetAlbum(*Token, "2UwqpfQtNuhBwviIC0f2ie") //Dasmso Spotify ID: 2UwqpfQtNuhBwviIC0f2ie
	if A.Error != "" {
		fmt.Println("Erreur lors de la récupération de l'album : ", A.Error, " ", A.ErrorDescription)
	} else {
		fmt.Println("\nAlbum récupéré : ", A.AlbumData)
		for i, item := range A.AlbumData {
			fmt.Printf("%d Nom de l'album: %s\nDate de sortie: %s\nNombre de pistes: %d\nURL Spotify: %s\nImage: %s\n\n",
				i, item.Name, item.ReleaseDate, item.TotalTracks, item.URL.Spotify, item.Image[1].URL)
		}
	}
	data := PageData{
		Title:   "Damso",
		Message: "Bienvenue sur la page de Damso 🎤",
	}
	tmpl := template.Must(template.ParseFiles("template/damso.html"))
	tmpl.Execute(w, data)
}

func Laylow(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Laylow",
		Message: "Bienvenue sur la page de Laylow 🎤",
	}
	tmpl := template.Must(template.ParseFiles("template/laylow.html"))
	tmpl.Execute(w, data)
}

// Contact gère la page de contact
func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // Si le formulaire est soumis en POST
		// Récupération des données du formulaire
		name := r.FormValue("name") // Récupère le champ "name"
		msg := r.FormValue("msg")   // Récupère le champ "msg"

		data := map[string]string{
			"Title":   "Contact",
			"Message": "Merci " + name + " pour ton message : " + msg, // Message personnalisé après soumission
		}
		renderTemplate(w, "contact.html", data)
		return // On termine ici pour ne pas exécuter la partie GET
	}

	// Si ce n'est pas un POST, on affiche simplement le formulaire
	data := map[string]string{
		"Title":   "Contact",
		"Message": "Envoie-nous un message 📩",
	}
	renderTemplate(w, "contact.html", data)
}
