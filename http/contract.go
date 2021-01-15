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

func Contract(w http.ResponseWriter, r *http.Request) {
	output := models.Contract()
	data := models.BookingData{}

	//Догрузка из БД
	conn, err := pgx.Connect(context.Background(), "postgres://unidb:Xz4lbm777_@45.90.34.227/hotel")
	help.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT " +
		"public.hotel.hotel_id, public.hotel.name, " +
		"public.city.city_id, public.city.name, " +
		"public.area.area_id, public.area.name, " +
		"public.room.room_id, public.room.room_number, " +
		"public.room_type.room_type_id, public.room_type.name, " +
		"public.room_info.price, public.room.status, " +
		"public.contract.contract_id, public.contract.arrival_date, public.contract.term, public.contract.target " +
		"FROM public.room " +
		"INNER JOIN public.room_info USING (room_info_id) " +
		"INNER JOIN public.hotel USING (hotel_id) " +
		"INNER JOIN public.room_type USING (room_type_id) " +
		"INNER JOIN public.city USING (city_id) " +
		"INNER JOIN public.area USING (area_id) " +
		"INNER JOIN public.contract USING (room_id) " +
		"ORDER BY public.contract.contract_id DESC")
	defer rows.Close()
	help.ErrDefaultDetect(err, "QueryRow")

	iData := []models.BookingData{}
	iCity:= []models.CityData{}
	iArea := []models.AreaData{}
	iType := []models.TypeData{}
	for rows.Next(){
		err = rows.Scan(&data.HotelId, &data.Hotel, &data.CityId, &data.City,
			&data.AreaId, &data.Area,&data.RoomId, &data.RoomNumber,
			&data.RoomTypeId, &data.RoomType, &data.Price, &data.Status,
			&data.ContactId, &data.ArrivalDate, &data.Term, &data.Target)
		data.Year, data.Month, data.Day = data.ArrivalDate.Date()
		iData = append(iData, data)
		flag := true
		for _, val := range iCity {
			if val.Id == data.CityId {
				flag = false
			}
		}
		if flag {
			iCity = append(iCity, models.CityData{data.CityId, data.City})
		}

		flag = true
		for _, val := range iArea {
			if val.Id == data.AreaId {
				flag = false
			}
		}
		if flag {
			iArea = append(iArea, models.AreaData{data.AreaId, data.Area})
		}

		flag = true
		for _, val := range iType {
			if val.Id == data.RoomTypeId {
				flag = false
			}
		}
		if flag {
			iType = append(iType, models.TypeData{data.RoomTypeId, data.RoomType})
		}

		helpers.ErrDefaultDetect(err, "Row Scan")
	}
	output.Booking = append(output.Booking, iData...)
	output.Cities = append(output.Cities, iCity...)
	output.Areas = append(output.Areas, iArea...)
	output.Types = append(output.Types, iType...)

	var fm = template.FuncMap{
		"mul": func(a, b int) int {
			return a * b
		},
	}

	templates := template.Must(template.New("contract").Funcs(fm).ParseFiles(
		"templates/header.tmpl",
		"templates/contract.tmpl",
		"templates/footer.tmpl"))

	tmpl := templates.Lookup("contract.tmpl")
	err = tmpl.ExecuteTemplate(w, "contract", output)

	help.ErrCatch(err, "Перевод в шаблон")
}
