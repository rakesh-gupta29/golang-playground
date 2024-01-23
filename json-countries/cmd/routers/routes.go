package routers

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
)


func MountAllRoutes() http.Handler {
  router := httprouter.New()
  router.ServeFiles("/static/*filepath",  http.Dir("public"))
  mountWebRoutes  ( router )
  mountAPIRoutes  ( router )
  return router 
}
