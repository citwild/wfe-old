package controllers

/**
 * Handles the front pages of the application
 */
import (
  "net/http"
  "path"
  "html/template"
)

func MainController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  //w.Write([]byte("This is a test"))
  filepath := path.Join("client/templates", "index.tmpl")
  tmpl, err := template.ParseFiles(filepath)
  forwardUserToTemplate(tmpl, w, err)
}

func ContactController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  filepath := path.Join("client/templates", "contact.tmpl")
  tmpl, err := template.ParseFiles(filepath)
  forwardUserToTemplate(tmpl, w, err)
}