package http_responses

import "net/http"

func ResponseBadRequest(w http.ResponseWriter, r *http.Request, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

// func ResponseBadRequest(w http.ResponseWriter, r *http.Request, msg string) {}
