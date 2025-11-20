package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ApiData struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Status  string   `json:"status"`
	Species string   `json:"species"`
	Type    string   `json:"type"`
	Gender  string   `json:"gender"`
	Image   string   `json:"image"`
	Episode []string `json:"episode"`
}

func GetToken() {

	// URL de L'API
	urlApi := "https://accounts.spotify.com/api/token"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	//Paramètres à inserer à la req POST
	data := url.Values{}
	data.Set("grant_type", "client_credentials") //Name, Value
	data.Set("client_id", "967f549670ed4aefbfedf2b746202c75")
	data.Set("client_secret", "ad8b9b61687d4c3387784e0a7e34c54e")

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodPost, urlApi, strings.NewReader(data.Encode())) // Méthode de req, url de l'API, Paramètres de la req (On converti les strings en flux lisible io.reader)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData ApiData

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	fmt.Println(decodeData)
}

func GetArtist() {

}

func GetAlbum() {

}

func GetTrack() {

}

func GetAPIdata() {

	// URL de L'API
	urlApi := "https://rickandmortyapi.com/api/character/120"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header, User_Agent permet d'identifier l'application, système ....
	req.Header.Add("User-Agent", "Ynov Campus Cours")

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData ApiData

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	fmt.Println(decodeData)
}
