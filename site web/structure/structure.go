package structure

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

type Token struct {
	AccessToken      string `json:"access_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// structure avec les données de l'album que l'on veut récupérer
type AllAlbums struct {
	AlbumData        []Items `json:"items"`
	Error            string  `json:"error"`
	ErrorDescription string  `json:"error_description"`
}

type Items struct {
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

type Track struct {
	Name             string   `json:"name"`
	Album            Album    `json:"album"`
	Artists          []Artist `json:"artists"`
	Error            string   `json:"error"`
	ErrorDescription string   `json:"error_description"`
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
