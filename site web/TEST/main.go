package main

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
	Token            string `json:"access_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type AlbumData struct {
	Data []Data
}
type Data struct {
	Image       string
	Name        string
	ReleaseDate string
	TotalTracks int
	URL         string
}

var Token *string

func main() {
	To := GetToken()
	if To.Error != "" {
		fmt.Println("Erreur lors de la récupération du token : ", To.Error, " ", To.ErrorDescription)
	} else {
		Token = &To.Token
		fmt.Println("Token récupéré : ", *Token)
	}

	A := GetAlbum(*Token, "2UwqpfQtNuhBwviIC0f2ie") //Dasmso ID: 2UwqpfQtNuhBwviIC0f2ie
	if A.Error != "" {
		fmt.Println("Erreur lors de la récupération de l'album : ", A.Error, " ", A.ErrorDescription)
	} else {
		fmt.Println("\nAlbum récupéré : ", A.AlbumItems)
		for i, item := range A.AlbumItems {
			fmt.Printf("%d Nom de l'album: %s\nDate de sortie: %s\nNombre de pistes: %d\nURL Spotify: %s\nImage: %s\n\n",
				i, item.Name, item.ReleaseDate, item.TotalTracks, item.URL.Spotify, item.Image[1].URL)
		}
	}

	Tr := GetTrack(*Token, "67Pf31pl0PfjBfUmvYNDCL") //Laylow Track ID: 3Gm5Z8u0U1J1R7aB2VY3fK
	if Tr.Error.Message != "" {
		fmt.Println("Erreur lors de la récupération du track : ", Tr.Error.Status, " ", Tr.Error.Message)
	} else {
		fmt.Printf("\nTrack récupéré : %s\nAlbum: %s\n", Tr.Name, Tr.Album.Name)
		var imgURL string
		if len(Tr.Album.Image) > 0 {
			imgURL = Tr.Album.Image[0].URL
		}
		fmt.Printf("%d Nom de musique: %s\nNom de l'artiste: %s\nNom de l'album: %s\nDate de sortie: %s\nURL Spotify: %s\nImage: %s\n\n",
			0, Tr.Name, Tr.Artists[0].Name, Tr.Album.Name, Tr.Album.ReleaseDate, Tr.Album.URL.Spotify, imgURL)
	}

	AHTML := AlbumData{}
	for _, i := range A.AlbumItems {
		data := Data{
			Image:       i.Image[1].URL,
			Name:        i.Name,
			ReleaseDate: i.ReleaseDate,
			TotalTracks: i.TotalTracks,
			URL:         i.URL.Spotify,
		}
		AHTML.Data = append(AHTML.Data, data)
	}
	fmt.Println("\n\n", AHTML)
	for i, d := range AHTML.Data {
		fmt.Printf("%d Image: %s\nNom de l'album: %s\nDate de sortie: %s\nTotal tracks: %d\nURL Spotify: %s\n\n", i, d.Image, d.Name, d.ReleaseDate, d.TotalTracks, d.URL)
	}

}

func GetToken() ApiData {
	// URL de L'API
	urlApi := "https://accounts.spotify.com/api/token"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	//Paramètres de type body à inserer à la req POST
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
		return ApiData{Error: errResp.Error()}
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
	if decodeData.Error != "" {
		return decodeData
	} else {
		fmt.Println("Token récupéré avec succès : ", decodeData.Token)
		return decodeData
	}
}

// structure avec les données de l'album que l'on veut récupérer
type AllAlbums struct {
	AlbumItems       []items `json:"items"`
	Error            string  `json:"error"`
	ErrorDescription string  `json:"error_description"`
}

type items struct {
	TotalTracks int         `json:"total_tracks"`
	URL         ExternalURL `json:"external_urls"`
	Image       []Image     `json:"images"`
	Name        string      `json:"name"`
	ReleaseDate string      `json:"release_date"`
}

type ExternalURL struct {
	Spotify string `json:"spotify"`
}

type Image struct {
	URL string `json:"url"`
}

func GetAlbum(Token string, id string) AllAlbums {
	// URL de L'API
	urlApi := "https://api.spotify.com/v1/artists/" + id + "/albums"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	//Paramètre de type query à inserer à la req GET
	q := req.URL.Query()
	q.Add("include_groups", "album")
	req.URL.RawQuery = q.Encode()

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+Token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return AllAlbums{Error: errResp.Error()}
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
	var decodeData AllAlbums

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	if decodeData.Error != "" {
		return decodeData
	} else {
		fmt.Printf("Album de %s récupéré : \n", id)
		println(decodeData.AlbumItems)
		return decodeData
	}
}

type Track struct {
	Name    string   `json:"name"`
	Album   Album    `json:"album"`
	Artists []Artist `json:"artists"`
	Error   Error    `json:"error"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Album struct {
	Name        string      `json:"name"`
	URL         ExternalURL `json:"external_urls"`
	Image       []Image     `json:"images"`
	ReleaseDate string      `json:"release_date"`
}

type Artist struct {
	Name string `json:"name"`
}

func GetTrack(Token string, id string) Track {
	// URL de L'API
	urlApi := "https://api.spotify.com/v1/tracks/" + id

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header
	req.Header.Add("Authorization", "Bearer "+Token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return Track{Error: Error{Message: errResp.Error()}}
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
	var decodeData Track

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	fmt.Println(string(body))

	// Affichage des données
	if decodeData.Error.Message != "" {
		return decodeData
	} else {
		fmt.Printf("Track de %s récupéré : \n", id)
		fmt.Println(decodeData)
		return decodeData
	}
}

func GetArtist() {

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
