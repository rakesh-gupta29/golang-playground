package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rakesh-gupta29/json-countries/cmd/handlers"
)

func mountAPIRoutes(router *httprouter.Router) *httprouter.Router {
	router.HandlerFunc(http.MethodGet, "/v1/all", handlers.GetAllCountries)
	router.HandlerFunc(http.MethodGet, "/v1/name/:name", handlers.GetMatchingCountries)
	router.HandlerFunc(http.MethodGet, "/v1/currency/:name", handlers.GetByCurrency)
	router.HandlerFunc(http.MethodGet, "/v1/demonym/:dem", handlers.GetByDemonym)
	router.HandlerFunc(http.MethodGet, "/v1/language/:lang", handlers.GetByLanguage)
	router.HandlerFunc(http.MethodGet, "/v1/capital/:capital", handlers.GetByCapital)
	router.HandlerFunc(http.MethodGet, "/v1/region/:region", handlers.GetByRegion)
	router.HandlerFunc(http.MethodGet, "/v1/subregion/:subregion", handlers.GetBySubregion)
	router.HandlerFunc(http.MethodGet, "/v1/translations/:name", handlers.GetbyTranslation)
	router.HandlerFunc(http.MethodGet, "/v1/filters/", handlers.GetByFilters)
	return router
}
