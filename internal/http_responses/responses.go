package http_responses

import "net/http"

func ResponseWithPayload(w http.ResponseWriter, r *http.Request, payload []byte) {
	w.Write(payload)
}