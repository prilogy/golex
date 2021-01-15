package models

import (
	s "course/models/static"
	"time"
)

type ContractData struct {
	Header 	s.HeaderData
	Booking	[]BookingData
	Cities 	[]CityData
	Areas  	[]AreaData
	Types  	[]TypeData
}

type BookingData struct {
	HotelId		int
	Hotel		string
	CityId		int
	City		string
	AreaId		int
	Area 		string
	RoomId 		int
	RoomNumber 	int
	RoomTypeId	int
	RoomType	string
	Price		int
	Status		bool
	ContactId	int
	ArrivalDate	time.Time
	Year		int
	Month		time.Month
	Day			int
	Term		int
	Target		string
}

func Contract() ContractData{
	data := ContractData{
		Header: s.Header(),
	}
	return data
}