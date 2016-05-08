// this is a test

package routers

import (
  "github.com/wide-field-ethnography/wfe/server/controllers"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
)

func SetMainRoutes( router *mux.Router) *mux.Router {

  // Front/Main page
  router.Handle("/", negroni.New(
    negroni.HandlerFunc(controllers.MainController),
  )).Methods("GET")

  // Contact Us page
  router.Handle("/contact", negroni.New(
    negroni.HandlerFunc(controllers.ContactController),
  )).Methods("GET")

  // TODO: Request access

  return router
}