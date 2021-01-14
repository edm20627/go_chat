package main

import (
	"net/http"

	"github.com/edm20627/go_chat/data"
)

func login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w , nil, "login.layout", "public.navbar", "signup")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
