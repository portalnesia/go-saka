/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type PancaSuda struct {
	LebuKatiupAngin Data
	WisesaSegara    Data
	TunggakSemi     Data
	SatriaWibawa    Data
	SumurSinaba     Data
	SatriaWirang    Data
	BumiKapetak     Data
}

func newPancaSuda() PancaSuda {
	return PancaSuda{
		LebuKatiupAngin: Data{0, "Lebu Katiup Angin"},
		WisesaSegara:    Data{1, "Wisesa Segara"},
		TunggakSemi:     Data{2, "Tunggak Semi"},
		SatriaWibawa:    Data{3, "Satria Wibawa"},
		SumurSinaba:     Data{4, "Sumur Sinaba"},
		SatriaWirang:    Data{5, "Satria Wirang"},
		BumiKapetak:     Data{6, "Bumi Kapetak"},
	}
}

func newPancaSudaSlice() []Data {
	data := newPancaSuda()
	return []Data{data.LebuKatiupAngin, data.WisesaSegara, data.TunggakSemi, data.SatriaWibawa, data.SumurSinaba, data.SatriaWirang, data.BumiKapetak}
}
