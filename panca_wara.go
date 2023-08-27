/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type PancaWaraData struct {
	ID    int64
	Urip  int64
	Kupih int64
	Name  string
}

type PancaWara struct {
	Paing  PancaWaraData
	Pon    PancaWaraData
	Wage   PancaWaraData
	Kliwon PancaWaraData
	Umanis PancaWaraData
}

func newPancaWara() PancaWara {
	return PancaWara{
		Paing:  PancaWaraData{0, 9, 3, "Paing"},
		Pon:    PancaWaraData{1, 7, 4, "Pon"},
		Wage:   PancaWaraData{2, 4, 5, "Wage"},
		Kliwon: PancaWaraData{3, 8, 1, "Kliwon"},
		Umanis: PancaWaraData{4, 5, 2, "Umanis"},
	}
}

func newPancaWaraSlice() []PancaWaraData {
	data := newPancaWara()
	return []PancaWaraData{data.Paing, data.Pon, data.Wage, data.Kliwon, data.Umanis}
}
