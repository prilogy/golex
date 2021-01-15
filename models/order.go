package models

import (
	s "course/models/static"
)

type OrderData struct{
	Header		s.HeaderData
	GuestInfo	[] FullGuestData
}

func Order() OrderData{
	data := OrderData{
		Header: s.Header(),
	}
	return data
}