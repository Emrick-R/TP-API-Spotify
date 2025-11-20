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

type Album struct { //structure avefs les données de l'album que l'on veut récupérer
	Items            []items `json:"items"`
	Error            string  `json:"error"`
	ErrorDescription string  `json:"error_description"`
}

type items struct {
	TotalTracks int         `json:"total_tracks"`
	URL         ExternalURL `json:"external_urls"`
	Image       Image       `json:"images"`
	Name        string      `json:"name"`
	ReleseDate  string      `json:"release_date"`
}

type ExternalURL struct {
	Spotify string `json:"spotify"`
}

type Image struct {
	URL string `json:"url"`
}
