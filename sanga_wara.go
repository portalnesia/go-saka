/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type SangaWara struct {
	Dangu   Wara
	Jangur  Wara
	Gigis   Wara
	Nohan   Wara
	Ogan    Wara
	Erangan Wara
	Urungan Wara
	Tulus   Wara
	Dadi    Wara
}

func newSangaWara() SangaWara {
	return SangaWara{
		Dangu:   Wara{0, 5, "Dangu"},
		Jangur:  Wara{1, 8, "Jangur"},
		Gigis:   Wara{2, 9, "Gigis"},
		Nohan:   Wara{3, 3, "Nohan"},
		Ogan:    Wara{4, 7, "Ogan"},
		Erangan: Wara{5, 1, "Erangan"},
		Urungan: Wara{6, 4, "Urungan"},
		Tulus:   Wara{7, 6, "Tulus"},
		Dadi:    Wara{8, 8, "Dadi"},
	}
}

func newSangaWaraSlice() []Wara {
	data := newSangaWara()
	return []Wara{
		data.Dangu,
		data.Jangur,
		data.Gigis,
		data.Nohan,
		data.Ogan,
		data.Erangan,
		data.Urungan,
		data.Tulus,
		data.Dadi,
	}
}
