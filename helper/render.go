package GroupieTracker

import (
	"net/http"
	"text/template"
)

type message_error struct {
	statut  int
	message string
}

func RenderError(w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles("tamplate/error.html")
	if err != nil {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(status)
	ErrorText := ""
	switch status {
	case 405:
		ErrorText = "YOU ARE NOT ALLOWED TO ENTER TO THIS PAGE."
	case 404:
		ErrorText = "THIS PAGE IS NOT FOUND"
	case 400:
		ErrorText = "BAD REQUEST"
	default:
		ErrorText = "INTERNAL SERVER ERROR"
	}
	pas_msg := message_error{statut: status, message: ErrorText}
	err = tmp.Execute(w, pas_msg)
	if err != nil {
		RenderError(w, http.StatusInternalServerError)
		return
	}
}
