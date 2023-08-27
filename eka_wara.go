/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type EkaWara struct {
	Void  Wara
	Luang Wara
}

func newEkaWara() EkaWara {
	return EkaWara{
		Void:  Wara{0, 0, ""},
		Luang: Wara{1, 1, "Luang"},
	}
}

func newEkaWaraSlice() []Wara {
	data := newEkaWara()
	return []Wara{data.Void, data.Luang}
}
