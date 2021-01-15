package http

import (
	"context"
	help "course/helpers"
	"course/models"
	"github.com/jackc/pgx"
	"html/template"
	"net/http"
	"servPanel/helpers"
)

func Guest(w http.ResponseWriter, r *http.Request) {
	output := models.Guest()
	data := models.FullGuestData{}

	//Догрузка из БД
	conn, err := pgx.Connect(context.Background(), "postgres://unidb:Xz4lbm777_@45.90.34.227/hotel")
	help.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM public.guest " +
		"ORDER BY public.guest.guest_id ASC")
	defer rows.Close()
	help.ErrDefaultDetect(err, "QueryRow")

	iData := []models.FullGuestData{}
	for rows.Next(){
		err = rows.Scan(&data.Id, &data.Fio, &data.Passport, &data.Birthday)
		data.Year, data.Month, data.Day = data.Birthday.Time.Date()
		iData = append(iData, data)
		helpers.ErrDefaultDetect(err, "Row Scan")
	}
	output.GuestInfo = append(output.GuestInfo, iData...)

	var fm = template.FuncMap{
		"mul": func(a, b int) int {
			return a * b
		},
	}

	templates := template.Must(template.New("guest").Funcs(fm).ParseFiles(
		"templates/header.tmpl",
		"templates/guest.tmpl",
		"templates/footer.tmpl"))

	tmpl := templates.Lookup("guest.tmpl")
	err = tmpl.ExecuteTemplate(w, "guest", output)

	help.ErrCatch(err, "Перевод в шаблон")
}
