/*
Copyright © Portalnesia <support@portalnesia.com>
*/

package saka

type Jejapan struct {
	Mina  Data
	Taru  Data
	Sato  Data
	Patra Data
	Wong  Data
	Paksi Data
}

func newJejapan() Jejapan {
	return Jejapan{
		Mina:  Data{0, "Mina"},
		Taru:  Data{1, "Taru"},
		Sato:  Data{2, "Sato"},
		Patra: Data{3, "Patra"},
		Wong:  Data{4, "Wong"},
		Paksi: Data{5, "Paksi"},
	}
}

func newJejapanSlice() []Data {
	data := newJejapan()
	return []Data{data.Mina, data.Taru, data.Sato, data.Patra, data.Wong, data.Paksi}
}
