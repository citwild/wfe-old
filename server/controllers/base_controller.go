package controllers

import (
  "html/template"
  "net/http"
)

func forwardUserToTemplate(tmpl *template.Template, w http.ResponseWriter, err error) {
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}


func createJwtCookie(token string) *http.Cookie {
  cookie := http.Cookie{}

  cookie.Name = "token"
  cookie.Value = token
  cookie.MaxAge = 1800 // 30 minutes

  return &cookie
}