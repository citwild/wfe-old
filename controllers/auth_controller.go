package controllers

import (
  "github.com/wide-field-ethnography/wfe/services"
  "github.com/wide-field-ethnography/wfe/services/models"
  "encoding/json"
  "net/http"
  "code.google.com/p/go-uuid/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {
  requestUser := new(models.User)

  requestUser.UUID = uuid.New()
  requestUser.Username = r.FormValue("username")
  requestUser.Password = r.FormValue("password")

  responseStatus, token := services.Login(requestUser)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(responseStatus)
  w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  requestUser := new(models.User)
  decoder     := json.NewDecoder(r.Body)

  decoder.Decode(&requestUser)

  w.Header().Set("Content-Type", "application/json")
  w.Write(services.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  err := services.Logout(r)
  w.Header().Set("Content-Type", "application/json")

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
  } else {
    w.WriteHeader(http.StatusOK)
  }
}