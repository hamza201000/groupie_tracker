package GroupieTracker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artists struct {
	Id              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationsUrl    string   `json:"locations"`
	CreationDateUrl int      `json:"creationDate"`
	RelationsUrl    string   `json:"relations"`

	Locations    Locations
	CreationDate Dates
	Relations    Relations
}
type Locations struct {
	Id       int      `json:"id"`
	Location []string `json:"location"`
	Dates    string   `json:"dates"`
}

type Dates struct {
	Id   int      `json:"id"`
	Date []string `json:"dates"`
}
type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetData() ([]Artists, error) {
	data, err := FitchData("https://groupietrackers.herokuapp.com/api/artists")
	return data, err
}

func FitchData(url string) ([]Artists, error) {
	var data []Artists
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}
	return &data, err
}
