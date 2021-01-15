package static

type HeaderData struct {
	NavBar []Menu
	Links  []Link
}

type Menu struct {
	Name string
	Link string
	Rate int
}

type Link struct {
	Name string
	Rel  string
	Href string
}

func Header() HeaderData {
	data := HeaderData{
		Links: []Link{
			{
				Name: "Главные стили",
				Rel:  "stylesheet",
				Href: "src/css/style.css",
			}},
		NavBar: []Menu{
			{
				Name: "Все номера",
				Link: "/",
				Rate: 1,
			},
			{
				Name: "Свободные номера",
				Link: "/",
				Rate: 2,
			},
			{
				Name: "Договоры",
				Link: "/contract",
				Rate: 3,
			},
			{
				Name: "Гости",
				Link: "/guest",
				Rate: 4,
			},
		},
	}

	return data
}
