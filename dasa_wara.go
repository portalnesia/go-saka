/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type DasaWara struct {
	Pandita Wara
	Pati    Wara
	Suka    Wara
	Duka    Wara
	Sri     Wara
	Manuh   Wara
	Manusa  Wara
	Raja    Wara
	Dewa    Wara
	Raksasa Wara
}

func newDasaWara() DasaWara {
	return DasaWara{
		Pandita: Wara{0, 5, "Pandita"},
		Pati:    Wara{1, 7, "Pati"},
		Suka:    Wara{2, 10, "Suka"},
		Duka:    Wara{3, 4, "Duka"},
		Sri:     Wara{4, 6, "Sri"},
		Manuh:   Wara{5, 2, "Manuh"},
		Manusa:  Wara{6, 3, "Manusa"},
		Raja:    Wara{7, 8, "Raja"},
		Dewa:    Wara{8, 9, "Dewa"},
		Raksasa: Wara{9, 1, "Raksasa"},
	}
}

func newDasaWaraSlice() []Wara {
	wara := newDasaWara()
	return []Wara{
		wara.Pandita,
		wara.Pati,
		wara.Suka,
		wara.Duka,
		wara.Sri,
		wara.Manuh,
		wara.Manusa,
		wara.Raja,
		wara.Dewa,
		wara.Raksasa,
	}
}
