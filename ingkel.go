/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type Ingkel struct {
	Wong  Data
	Sato  Data
	Mina  Data
	Manuk Data
	Taru  Data
	Buku  Data
}

func newIngkel() Ingkel {
	return Ingkel{
		Wong:  Data{0, "Wong"},
		Sato:  Data{1, "Sato"},
		Mina:  Data{2, "Mina"},
		Manuk: Data{3, "Manuk"},
		Taru:  Data{4, "Taru"},
		Buku:  Data{5, "Buku"},
	}
}

func newIngkelSlice() []Data {
	data := newIngkel()
	return []Data{data.Wong, data.Sato, data.Mina, data.Manuk, data.Taru, data.Buku}
}
