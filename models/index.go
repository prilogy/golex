package models

import (
	s "course/models/static"
)

type ViewData struct {
	Header s.HeaderData
	Room   []RoomData
	Cities []CityData
	Areas  []AreaData
	Types  []TypeData
}

type CityData struct {
	Id		int
	Name 	string
}
type AreaData struct {
	Id		int
	Name 	string
}
type TypeData struct {
	Id		int
	Name 	string
}


type RoomData struct {
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
}

func Index() ViewData{
	data := ViewData{
		Header: s.Header(),
	}
	return data
}