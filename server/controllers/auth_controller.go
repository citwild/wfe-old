package controllers

import (
  "github.com/wide-field-ethnography/wfe/server/services"
  "github.com/wide-field-ethnography/wfe/server/services/models"
  "encoding/json"
  "net/http"
  "code.google.com/p/go-uuid/uuid"
  "path"
  "html/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
  requestUser := new(models.User)

  requestUser.UUID = uuid.New()
  requestUser.Username = r.FormValue("username")
  requestUser.Password = r.FormValue("password")

  responseStatus, token := services.Login(requestUser)

  // authentication successful, direct to landing page
  if responseStatus == http.StatusOK {
    filepath := path.Join("client/templates", "bucketlist.tmpl")

    //provide token in cookie (bad practice?)
    http.SetCookie(w, createJwtCookie(string(token)))
    w.WriteHeader(responseStatus)

    tmpl, err := template.ParseFiles(filepath)
    forwardUserToTemplate(tmpl, w, err)
  }
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