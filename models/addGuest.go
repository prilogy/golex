package models

import (
	s "course/models/static"
	"github.com/jackc/pgtype"
	"time"
)

type AddGuestData struct{
	Header		s.HeaderData
	Guest		[] FullGuestData
}

type FullGuestData struct {
	Id			int
	Fio 		string
	Passport 	int
	Birthday	pgtype.Date
	Year		int
	Month		time.Month
	Day			int
}

func AddGuest() AddGuestData{
	data := AddGuestData{
		Header: s.Header(),
	}
	return data
}