package module

type Artists struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	DatesLocations map[string][]string
}

var (
	Module      []Artists
	relation    Relations
	artisturl   string = "https://groupietrackers.herokuapp.com/api/artists"
	relationurl string = "https://groupietrackers.herokuapp.com/api/relation"
)

type AllArtists struct {
	Artists []*Artists
}
type Order struct {
	
}
type Relations struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
