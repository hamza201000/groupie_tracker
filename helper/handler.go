package GroupieTracker

import (
	"net/http"
	"os"
	"text/template"
)



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

	data_artist, err := GetData()
	if err != nil {
		RenderError(w, http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, data_artist)
	if err != nil {
		RenderError(w, http.StatusInternalServerError)
		return
	}
}

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed)
		return
	}
	path, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		RenderError(w, http.StatusInternalServerError)
		return
	} else if path.IsDir() {
		RenderError(w, http.StatusNotFound)
		return
	} else {
		http.ServeFile(w, r, r.URL.Path[1:])
	}
}
