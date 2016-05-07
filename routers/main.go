// this is a test

package routers

import (
  "github.com/wide-field-ethnography/wfe/controllers"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
)

func SetMainRoutes( router *mux.Router) *mux.Router {
  router.Handle("/", negroni.New(
    negroni.HandlerFunc(controllers.MainController),
  )).Methods("GET")

  return router
}