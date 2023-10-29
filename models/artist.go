package models

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relation struct {
	Id            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}

type RelationsData struct {
	Index []Relation `json:"index"`
}

var (
	Artists   []Artist
	Relations RelationsData
)

const (
	ArtistsURL  = "https://groupietrackers.herokuapp.com/api/artists"
	RelationURL = "https://groupietrackers.herokuapp.com/api/relation"
)
