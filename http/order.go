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

func Order(w http.ResponseWriter, r *http.Request) {
	output := models.Order()
	data := models.FullGuestData{}

	//Догрузка из БД
	conn, err := pgx.Connect(context.Background(), "postgres://unidb:Xz4lbm777_@45.90.34.227/hotel")
	help.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT " +
		"public.guest.guest_id, public.guest.fio FROM public.guest ")
	defer rows.Close()
	help.ErrDefaultDetect(err, "QueryRow")

	iData := []models.FullGuestData{}
	for rows.Next(){
		err = rows.Scan(&data.Id, &data.Fio)
		iData = append(iData, data)
		helpers.ErrDefaultDetect(err, "Row Scan")
	}

	output.GuestInfo = append(output.GuestInfo, iData...)

	templates, err := template.ParseFiles(
		"templates/header.tmpl",
		"templates/order.tmpl",
		"templates/footer.tmpl")

	tmpl := templates.Lookup("order.tmpl")
	err = tmpl.ExecuteTemplate(w, "order", output)

	help.ErrCatch(err, "Перевод в шаблон")
}
