package handler

import (
	"groupie-tracker/internal/models"
	"net/http"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, code int, title, message string) {
	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("web/error.html")
	if err != nil {
		http.Error(w, http.StatusText(code), code)
		return
	}
	data := models.ErrorPage{
		Code:    code,
		Title:   title,
		Message: message,
	}

	tmpl.Execute(w, data)

}
