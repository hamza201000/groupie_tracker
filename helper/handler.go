package GroupieTracker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

type data_artist struct {
	artists   string `json:"artists"`
	locations string `json:"locations"`
	dates     string `json:"dates"`
	relation  string `json:"relation"`
}
type data_locations struct {
	id       int      `json:"id"`
	location []string `json:"location"`
}

type data_dates struct {
	id    int      `json:"id"`
	dates []string `json:"dates"`
}
type relation struct {
	id             int                 `json:"id"`
	datesLocations map[string][]string `json:"datesLocations"`
}

type data struct {
	artists   string
	locations string
	dates     string
	relation  string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("tamplate/index.html")
	if err != nil {
		RenderError(w, http.StatusInternalServerError)
		return
	}
	data_artist := data{artists: GetData().artists, locations: GetData().locations, dates: GetData().dates, relation: GetData().relation}
	err = tmp.Execute(w, data_artist)
	if err != nil {
		RenderError(w, http.StatusInternalServerError)
		return
	}
}

func GetData() data_artist {
	var data data_artist
	url := "https://groupietrackers.herokuapp.com/api"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return data
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return data
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	return data
}
