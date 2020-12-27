package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/edm20627/go_chat/data"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []stirng
	for _, file := range filenames {
		files = append(files, fmt.Fprintf("template/%s.html", file))
	}

	templates := template.Must(template.ParseFile(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
