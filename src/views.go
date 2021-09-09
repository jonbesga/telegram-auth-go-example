package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates/*
var templates embed.FS

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFS(templates, "templates/index.html")
	t.Execute(w, nil)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFS(templates, "templates/profile.html")
	t.Execute(w, nil)
}

func AuthTelegramHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	first_name := query.Get("first_name")
	username := query.Get("username")
	photo_url := query.Get("photo_url")
	auth_date := query.Get("auth_date")
	hash := query.Get("hash")

	data_check_string := []byte(fmt.Sprintf("auth_date=%s\nfirst_name=%s\nid=%s\nphoto_url=%s\nusername=%s", auth_date, first_name, id, photo_url, username))

	isFromTelegram := VerifyTelegramData(data_check_string, hash)
	if isFromTelegram {
		cookie := http.Cookie{Name: "token", Value: hash, Path: "/"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	} else {
		w.Write([]byte("FAIL"))
	}
}
