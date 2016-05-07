package controllers

import (
  "net/http"
)

func MainController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  w.Write([]byte("This is a test"))
}