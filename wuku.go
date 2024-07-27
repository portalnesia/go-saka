/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type WukuData struct {
	ID   int64
	Urip int64
	Name string
}

type Wuku struct {
	Sinta        WukuData
	Landep       WukuData
	Ukir         WukuData
	Kulantir     WukuData
	Tolu         WukuData
	Gumbreg      WukuData
	Wariga       WukuData
	Warigadean   WukuData
	Julungwangi  WukuData
	Sungsang     WukuData
	Dungulan     WukuData
	Kuningan     WukuData
	Langkir      WukuData
	Medangsia    WukuData
	Pujut        WukuData
	Pahang       WukuData
	Krulut       WukuData
	Merakih      WukuData
	Tambir       WukuData
	Medangkungan WukuData
	Matal        WukuData
	Uye          WukuData
	Menail       WukuData
	Prangbakat   WukuData
	Bala         WukuData
	Ugu          WukuData
	Wayang       WukuData
	Klawu        WukuData
	Dukut        WukuData
	Watugunung   WukuData
}

func newWuku() Wuku {
	return Wuku{
		Sinta:        WukuData{0, 7, "Sinta"},
		Landep:       WukuData{1, 1, "Landep"},
		Ukir:         WukuData{2, 4, "Ukir"},
		Kulantir:     WukuData{3, 6, "Kulantir"},
		Tolu:         WukuData{4, 5, "Tolu"},
		Gumbreg:      WukuData{5, 8, "Gumbreg"},
		Wariga:       WukuData{6, 9, "Wariga"},
		Warigadean:   WukuData{7, 3, "Warigadean"},
		Julungwangi:  WukuData{8, 7, "Julungwangi"},
		Sungsang:     WukuData{9, 1, "Sungsang"},
		Dungulan:     WukuData{10, 4, "Dungulan"},
		Kuningan:     WukuData{11, 6, "Kuningan"},
		Langkir:      WukuData{12, 5, "Langkir"},
		Medangsia:    WukuData{13, 8, "Medangsia"},
		Pujut:        WukuData{14, 9, "Pujut"},
		Pahang:       WukuData{15, 3, "Pahang"},
		Krulut:       WukuData{16, 7, "Krulut"},
		Merakih:      WukuData{17, 1, "Merakih"},
		Tambir:       WukuData{18, 4, "Tambir"},
		Medangkungan: WukuData{19, 6, "Medangkungan"},
		Matal:        WukuData{20, 5, "Matal"},
		Uye:          WukuData{21, 8, "Uye"},
		Menail:       WukuData{22, 9, "Menail"},
		Prangbakat:   WukuData{23, 3, "Prangbakat"},
		Bala:         WukuData{24, 7, "Bala"},
		Ugu:          WukuData{25, 1, "Ugu"},
		Wayang:       WukuData{26, 4, "Wayang"},
		Klawu:        WukuData{27, 6, "Klawu"},
		Dukut:        WukuData{28, 5, "Dukut"},
		Watugunung:   WukuData{29, 8, "Watugunung"},
	}
}

func newWukuSlice() []WukuData {
	data := newWuku()
	return []WukuData{
		data.Sinta,
		data.Landep,
		data.Ukir,
		data.Kulantir,
		data.Tolu,
		data.Gumbreg,
		data.Wariga,
		data.Warigadean,
		data.Julungwangi,
		data.Sungsang,
		data.Dungulan,
		data.Kuningan,
		data.Langkir,
		data.Medangsia,
		data.Pujut,
		data.Pahang,
		data.Krulut,
		data.Merakih,
		data.Tambir,
		data.Medangkungan,
		data.Matal,
		data.Uye,
		data.Menail,
		data.Prangbakat,
		data.Bala,
		data.Ugu,
		data.Wayang,
		data.Klawu,
		data.Dukut,
		data.Watugunung,
	}
}
