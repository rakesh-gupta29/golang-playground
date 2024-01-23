package handlers

import (
	"net/http"

	"github.com/rakesh-gupta29/json-countries/pkg/res"
)

type Check struct {
	Name string `json:"name"`
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	books := []Check{{Name: "checking"}}
	res.WriteJSON(w, books, http.StatusOK, nil)
}
