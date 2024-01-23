
package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rakesh-gupta29/json-countries/cmd/handlers"
)

func mountWebRoutes(router *httprouter.Router) *httprouter.Router {
  router.HandlerFunc(http.MethodGet,"/", handlers.HandleHome) 
 return router 
}
