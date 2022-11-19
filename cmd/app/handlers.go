package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sudoku/internal/api"
	resp "sudoku/internal/http_responses"
)

func (app *App) SolveHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		msg := fmt.Sprint("cannot read request body:", err.Error())
		app.L.Error(msg)
		resp.ResponseBadRequest(w, r, msg)
		return
	}

	res, err := api.SolveBoard(req)
	if err != nil {
		app.L.Error(err.Error())
		resp.ResponseBadRequest(w, r, err.Error())
		return
	}

	resp.ResponseWithPayload(w, r, res)
}
