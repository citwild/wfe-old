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
  filepath := path.Join("templates", "index.tmpl")
  tmpl, err := template.ParseFiles(filepath)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func ContactController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  filepath := path.Join("templates", "contact.tmpl")
  tmpl, err := template.ParseFiles(filepath)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}