/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type SaptaWaraData struct {
	ID       int64
	Urip     int64
	KertaAji int64
	Kupih    int64
	Name     string
}

type SaptaWara struct {
	Redite    SaptaWaraData
	Soma      SaptaWaraData
	Anggara   SaptaWaraData
	Buda      SaptaWaraData
	Wraspati  SaptaWaraData
	Sukra     SaptaWaraData
	Saniscara SaptaWaraData
}

func newSaptaWara() SaptaWara {
	return SaptaWara{
		Redite:    SaptaWaraData{0, 5, 6, 3, "Redite"},
		Soma:      SaptaWaraData{1, 4, 4, 4, "Soma"},
		Anggara:   SaptaWaraData{2, 3, 3, 5, "Anggara"},
		Buda:      SaptaWaraData{3, 7, 6, 6, "Buda"},
		Wraspati:  SaptaWaraData{4, 8, 5, 7, "Wraspati"},
		Sukra:     SaptaWaraData{5, 6, 7, 1, "Sukra"},
		Saniscara: SaptaWaraData{6, 9, 8, 2, "Saniscara"},
	}
}

func newSaptaWaraSlice() []SaptaWaraData {
	data := newSaptaWara()
	return []SaptaWaraData{data.Redite, data.Soma, data.Anggara, data.Buda, data.Wraspati, data.Sukra, data.Saniscara}
}
