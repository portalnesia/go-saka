/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type PratithiSamutPada struct {
	Tresna     Data
	Upadana    Data
	Bhawa      Data
	Jati       Data
	Jaramarana Data
	Awidya     Data
	Saskara    Data
	Widnyana   Data
	Namarupa   Data
	Sadayatana Data
	Separsa    Data
	Wedana     Data
}

func newPratithiSamutPada() PratithiSamutPada {
	return PratithiSamutPada{
		Tresna:     Data{0, "Tresna"},
		Upadana:    Data{1, "Upadana"},
		Bhawa:      Data{2, "Bhawa"},
		Jati:       Data{3, "Jati"},
		Jaramarana: Data{4, "Jaramarana"},
		Awidya:     Data{5, "Awidya"},
		Saskara:    Data{6, "Saskara"},
		Widnyana:   Data{7, "Widnyana"},
		Namarupa:   Data{8, "Namarupa"},
		Sadayatana: Data{9, "Sadayatana"},
		Separsa:    Data{10, "Separsa"},
		Wedana:     Data{11, "Wedana"},
	}
}

func newPratithiSamutPadaSlice() []Data {
	data := newPratithiSamutPada()
	return []Data{
		data.Tresna,
		data.Upadana,
		data.Bhawa,
		data.Jati,
		data.Jaramarana,
		data.Awidya,
		data.Saskara,
		data.Widnyana,
		data.Namarupa,
		data.Sadayatana,
		data.Separsa,
		data.Wedana,
	}
}
