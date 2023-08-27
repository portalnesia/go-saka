/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type Lintang struct {
	Gajah        Data
	Kiriman      Data
	JungSarat    Data
	AtiwaTiwa    Data
	SangkaTikel  Data
	BubuBolong   Data
	Sungenge     Data
	Uluku        Data
	Pedati       Data
	Kuda         Data
	GajahMina    Data
	Bade         Data
	Magelut      Data
	Pagelangan   Data
	KalaSungsang Data
	Kukus        Data
	Asu          Data
	Kartika      Data
	Naga         Data
	AngsaAngrem  Data
	Panah        Data
	Patrem       Data
	Lembu        Data
	Depat        Data
	Tangis       Data
	SalahUkur    Data
	PerahuPegat  Data
	PuwuhAtarung Data
	Lawean       Data
	Kelapa       Data
	Yuyu         Data
	Lumbung      Data
	Kumba        Data
	Udang        Data
	Begoong      Data
}

func newLintang() Lintang {
	return Lintang{
		Gajah:        Data{0, "Gajah"},
		Kiriman:      Data{1, "Kiriman"},
		JungSarat:    Data{2, "Jung Sarat"},
		AtiwaTiwa:    Data{3, "Atiwa Tiwa"},
		SangkaTikel:  Data{4, "Sangka Tikel"},
		BubuBolong:   Data{5, "Bubu Bolong"},
		Sungenge:     Data{6, "Sungenge"},
		Uluku:        Data{7, "Uluku"},
		Pedati:       Data{8, "Pedati"},
		Kuda:         Data{9, "Kuda"},
		GajahMina:    Data{10, "Gajah Mina"},
		Bade:         Data{11, "Bade"},
		Magelut:      Data{12, "Magelut"},
		Pagelangan:   Data{13, "Pagelangan"},
		KalaSungsang: Data{14, "Kala Sungsang"},
		Kukus:        Data{15, "Kukus"},
		Asu:          Data{16, "Asu"},
		Kartika:      Data{17, "Kartika"},
		Naga:         Data{18, "Naga"},
		AngsaAngrem:  Data{19, "Angsa Angrem"},
		Panah:        Data{20, "Panah"},
		Patrem:       Data{21, "Patrem"},
		Lembu:        Data{22, "Lembu"},
		Depat:        Data{23, "Depat"},
		Tangis:       Data{24, "Tangis"},
		SalahUkur:    Data{25, "Salah Ukur"},
		PerahuPegat:  Data{26, "Perahu Pegat"},
		PuwuhAtarung: Data{27, "Puwuh Atarung"},
		Lawean:       Data{28, "Lawean"},
		Kelapa:       Data{29, "Kelapa"},
		Yuyu:         Data{30, "Yuyu"},
		Lumbung:      Data{31, "Lumbung"},
		Kumba:        Data{32, "Kumba"},
		Udang:        Data{33, "Udang"},
		Begoong:      Data{34, "Begoong"},
	}
}

func newLintangSlice() []Data {
	data := newLintang()
	return []Data{
		data.Gajah,
		data.Kiriman,
		data.JungSarat,
		data.AtiwaTiwa,
		data.SangkaTikel,
		data.BubuBolong,
		data.Sungenge,
		data.Uluku,
		data.Pedati,
		data.Kuda,
		data.GajahMina,
		data.Bade,
		data.Magelut,
		data.Pagelangan,
		data.KalaSungsang,
		data.Kukus,
		data.Asu,
		data.Kartika,
		data.Naga,
		data.AngsaAngrem,
		data.Panah,
		data.Patrem,
		data.Lembu,
		data.Depat,
		data.Tangis,
		data.SalahUkur,
		data.PerahuPegat,
		data.PuwuhAtarung,
		data.Lawean,
		data.Kelapa,
		data.Yuyu,
		data.Lumbung,
		data.Kumba,
		data.Udang,
		data.Begoong,
	}
}
