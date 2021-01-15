package models

import (
	s "course/models/static"
)

type AllGuestData struct{
	Header		s.HeaderData
	GuestInfo	[] FullGuestData
}

func Guest() AllGuestData{
	data := AllGuestData{
		Header: s.Header(),
	}
	return data
}