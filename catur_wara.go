/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type CaturWara struct {
	Sri    Wara
	Laba   Wara
	Jaya   Wara
	Menala Wara
}

func newCaturWara() CaturWara {
	return CaturWara{
		Sri:    Wara{0, 6, "Sri"},
		Laba:   Wara{1, 5, "Laba"},
		Jaya:   Wara{2, 1, "Jaya"},
		Menala: Wara{3, 8, "Menala"},
	}
}

func newCaturWaraSlice() []Wara {
	wara := newCaturWara()
	return []Wara{wara.Sri, wara.Laba, wara.Jaya, wara.Menala}
}
