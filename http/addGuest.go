package http

import (
	help "course/helpers"
	"course/models"
	"html/template"
	"net/http"
)

func AddGuest(w http.ResponseWriter, r *http.Request) {
	output := models.AddGuest()

	templates, err := template.ParseFiles(
		"templates/header.tmpl",
		"templates/addGuest.tmpl",
		"templates/footer.tmpl")

	tmpl := templates.Lookup("addGuest.tmpl")
	err = tmpl.ExecuteTemplate(w, "addGuest", output)

	help.ErrCatch(err, "Перевод в шаблон")
}
