package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rakesh-gupta29/json-countries/cmd/models"
	"github.com/rakesh-gupta29/json-countries/pkg/res"
)

var countries []models.Country

func init() {
	jsonFile, err := os.Open("assets/countries.json")
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error marshaling the file")
	}
	json.Unmarshal(byteData, &countries)
	defer jsonFile.Close()

}

func GetAllCountries(w http.ResponseWriter, r *http.Request) {
	res.WriteJSON(w, countries, http.StatusOK, nil)
}

// /v1/name/{name of the country}
func GetMatchingCountries(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	parts := strings.Split(path, "/")
	query := ""
	if len(parts) > 2 {
		query = parts[3]
	} else {
		query = ""
	}

	isFullName := r.URL.Query().Get("fullName")

	matchingCountries := []models.Country{}
	if isFullName == "true" {
		for _, elem := range countries {
			commonName := strings.ToLower(elem.Name.Common)
			officialName := strings.ToLower(elem.Name.Official)
			if commonName == query {
				matchingCountries = append(matchingCountries, elem)
			} else if officialName == query {
				matchingCountries = append(matchingCountries, elem)
			}
		}
	} else {
		for _, elem := range countries {
			if strings.Contains(strings.ToLower(elem.Name.Common), strings.ToLower(query)) {
				matchingCountries = append(matchingCountries, elem)
			} else if strings.Contains(strings.ToLower(elem.Name.Official), strings.ToLower(query)) {
				matchingCountries = append(matchingCountries, elem)
			} else {
				for _, names := range elem.Name.NativeName {
					if strings.Contains(strings.ToLower(names.Common), strings.ToLower(query)) {
						matchingCountries = append(matchingCountries, elem)
					}

				}
			}
		}
	}

	if len(matchingCountries) == 0 {
		res.WriteJSON(w, matchingCountries, http.StatusNotFound, nil)
		return
	}
	res.WriteJSON(w, matchingCountries, http.StatusOK, nil)
}

// v1/currency/{name of the currency or code for the currency}
func GetByCurrency(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	query := ""
	if len(parts) > 2 {
		query = parts[3]
	} else {
		query = ""
	}
	matchingCountries := []models.Country{}
	for _, elem := range countries {
		for key, _ := range elem.Currencies {
			if strings.EqualFold(strings.ToLower(key), strings.ToLower(query)) {
				matchingCountries = append(matchingCountries, elem)
			}
		}
	}

	if len(matchingCountries) == 0 {
		res.WriteJSON(w, matchingCountries, http.StatusNotFound, nil)
		return
	}
	res.WriteJSON(w, matchingCountries, http.StatusOK, nil)
}

// v1/demonym/{male or female in any language}
func GetByDemonym(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	query := ""
	if len(parts) > 2 {
		query = strings.ToLower(parts[3])
	} else {
		query = ""
	}

	matchingCountries := []models.Country{}

	for _, elem := range countries {
		for _, i := range elem.Demonyms {
			if strings.ToLower(i.F) == query || strings.ToLower(i.M) == query {
				matchingCountries = append(matchingCountries, elem)
			}
		}
	}

	if len(matchingCountries) == 0 {
		res.WriteJSON(w, matchingCountries, http.StatusNotFound, nil)
		return
	}
	res.WriteJSON(w, matchingCountries, http.StatusOK, nil)
}

// v1/langauges/{name or code}
func GetByLanguage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	query := ""
	if len(parts) > 2 {
		query = strings.ToLower(parts[3])
	} else {
		query = ""
	}

	matchingCountries := []models.Country{}

	for _, elem := range countries {
		for key, value := range elem.Languages {
			if strings.ToLower(key) == query || strings.ToLower(value) == query {
				matchingCountries = append(matchingCountries, elem)
			}

		}
	}

	if len(matchingCountries) == 0 {
		res.WriteJSON(w, matchingCountries, http.StatusNotFound, nil)
		return
	}
	res.WriteJSON(w, matchingCountries, http.StatusOK, nil)
}

// v1/capital/{capital city}
func GetByCapital(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	query := ""
	if len(parts) > 2 {
		query = strings.ToLower(parts[3])
	} else {
		query = ""
	}

	matchingCountries := []models.Country{}

	for _, elem := range countries {
		for _, value := range elem.Capital {
			if strings.Contains(strings.ToLower(value), query) {
				matchingCountries = append(matchingCountries, elem)
			}

		}
	}

	if len(matchingCountries) == 0 {
		res.WriteJSON(w, matchingCountries, http.StatusNotFound, nil)
		return
	}
	res.WriteJSON(w, matchingCountries, http.StatusOK, nil)
}

// v1/region/{region name or code}
func GetByRegion(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	query := ""
	if len(parts) > 2 {
		query = strings.ToLower(parts[3])
	} else {
		query = ""
	}
	matchingCountries := []models.Country{}

	for _, elem := range countries {
		if strings.Contains(strings.ToLower(elem.Region), query) {
			matchingCountries = append(matchingCountries, elem)
		}
	}

	if len(matchingCountries) == 0 {
		res.WriteJSON(w, matchingCountries, http.StatusNotFound, nil)
		return
	}
	res.WriteJSON(w, matchingCountries, http.StatusOK, nil)

}

// v1/subregion/{subregion name or code}
func GetBySubregion(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	query := ""
	if len(parts) > 2 {
		query = strings.ToLower(parts[3])
	} else {
		query = ""
	}
	matchingCountries := []models.Country{}

	for _, elem := range countries {
		if strings.Contains(strings.ToLower(elem.Subregion), query) {
			matchingCountries = append(matchingCountries, elem)
		}
	}

	if len(matchingCountries) == 0 {
		res.WriteJSON(w, matchingCountries, http.StatusNotFound, nil)
		return
	}
	res.WriteJSON(w, matchingCountries, http.StatusOK, nil)
}

// v1/translation/{any translation}
func GetbyTranslation(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	query := ""
	if len(parts) > 2 {
		query = strings.ToLower(parts[3])
	} else {
		query = ""
	}
	matchingCountries := []models.Country{}

	for _, elem := range countries {
		for _, value := range elem.Translations {
			if strings.Contains(strings.ToLower(value.Common), query) {
				matchingCountries = append(matchingCountries, elem)
			}
		}
	}

	if len(matchingCountries) == 0 {
		res.WriteJSON(w, matchingCountries, http.StatusNotFound, nil)
		return
	}
	res.WriteJSON(w, matchingCountries, http.StatusOK, nil)
}

// v1/filter/?key=value&key=value
type Query struct {
	UnMember    *bool `json:"unMember"`
	Landlocked  *bool `json:"landlocked"`
	Independent *bool `json:"independent"`
}

func isValidKey(key string) bool {
	validKeys := []string{
		"independent",
		"unMember",
		"landlocked",
	}

	for _, validKey := range validKeys {
		if key == validKey {
			return true
		}
	}

	return false
}

func boolPtr(b bool) *bool {
	return &b
}

func fetchQueryParams(r *http.Request) (*Query, error) {

	queryValues := r.URL.Query()

	req := &Query{}

	for key := range queryValues {
		if isValidKey(key) {
			value := queryValues.Get(key)
			switch key {
			case "independent":
				req.Independent = boolPtr((value == "true"))
			case "unMember":
				req.UnMember = boolPtr((value == "true"))
			case "landlocked":
				req.Landlocked = boolPtr((value == "true"))
			}
		} else {
			return nil, errors.New("Invalid key: " + key)
		}
	}

	return req, nil

}

func getMathingCountries(query *Query) []models.Country {

	var filtered []models.Country

	for _, elem := range countries {
		if (query.Independent == nil || *query.Independent == elem.Independent) &&
			(query.Landlocked == nil || *query.Landlocked == elem.Landlocked) &&
			(query.UnMember == nil || *query.UnMember == elem.UnMember) {
			filtered = append(filtered, elem)
		}
	}
	return filtered

}
func GetByFilters(w http.ResponseWriter, r *http.Request) {
	query, err := fetchQueryParams(r)
	if err != nil {
		resp := models.ValidationError{IsOK: false, Message: err.Error()}
		res.WriteJSON(w, resp, http.StatusBadRequest, nil)
		return
	}

	matches := getMathingCountries(query)
	res.WriteJSON(w, matches, http.StatusOK, nil)
}
