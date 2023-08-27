/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type Rahinan struct {
	SomaRibek          Data
	SabuhEmas          Data
	PagerWesi          Data
	TumpekLandep       Data
	TumpekUduh         Data
	SugihanJawa        Data
	SugihanBali        Data
	PenyekebanGalungan Data
	PenyajanGalungan   Data
	PenampahanGalungan Data
	Galungan           Data
	ManisGalungan      Data
	PemaridanGuru      Data
	Ulihan             Data
	PemacekanAgung     Data
	PenampahanKuningan Data
	Kuningan           Data
	PegatUwakan        Data
	TumpekKandang      Data
	TumpekWayang       Data
	Saraswati          Data
	BanyuPinaruh       Data
	SiwaRatri          Data
	TawurAgungKasanga  Data
	Nyepi              Data
	NgembakGeni        Data
	BudaCemeng         Data
	AnggaraKasih       Data
	KajengKliwon       Data
	Purnama            Data
	Tilem              Data
}

func newRahinan() Rahinan {
	return Rahinan{
		SomaRibek:          Data{0, "Soma Ribek"},
		SabuhEmas:          Data{1, "Sabuh Emas"},
		PagerWesi:          Data{2, "Pager Wesi"},
		TumpekLandep:       Data{3, "Tumpek Landep"},
		TumpekUduh:         Data{4, "Tumpek Uduh"},
		SugihanJawa:        Data{5, "Sugihan Jawa"},
		SugihanBali:        Data{6, "Sugihan Bali"},
		PenyekebanGalungan: Data{7, "Penyekeban Galungan"},
		PenyajanGalungan:   Data{8, "Penyajan Galungan"},
		PenampahanGalungan: Data{9, "Penampahan Galungan"},
		Galungan:           Data{10, "Galungan"},
		ManisGalungan:      Data{11, "Manis Galungan"},
		PemaridanGuru:      Data{12, "Pemaridan Guru"},
		Ulihan:             Data{13, "Ulihan"},
		PemacekanAgung:     Data{14, "Pemacekan Agung"},
		PenampahanKuningan: Data{15, "Penampahan Kuningan"},
		Kuningan:           Data{16, "Kuningan"},
		PegatUwakan:        Data{17, "Pegat Uwakan"},
		TumpekKandang:      Data{18, "Tumpek Kandang"},
		TumpekWayang:       Data{19, "Tumpek Wayang"},
		Saraswati:          Data{20, "Saraswati"},
		BanyuPinaruh:       Data{21, "Banyu Pinaruh"},
		SiwaRatri:          Data{22, "Siwa Ratri"},
		TawurAgungKasanga:  Data{23, "Tawur Agung Kasanga"},
		Nyepi:              Data{24, "Nyepi"},
		NgembakGeni:        Data{25, "Ngembak Geni"},
		BudaCemeng:         Data{26, "Buda Cemeng"},
		AnggaraKasih:       Data{27, "Anggara Kasih"},
		KajengKliwon:       Data{28, "Kajeng Kliwon"},
		Purnama:            Data{29, "Purnama"},
		Tilem:              Data{30, "Tilem"},
	}
}

// func newRahinanSlice() []Data {
// 	data := newRahinan()
// 	return []Data{
// 		data.SomaRibek,
// 		data.SabuhEmas,
// 		data.PagerWesi,
// 		data.TumpekLandep,
// 		data.TumpekUduh,
// 		data.SugihanJawa,
// 		data.SugihanBali,
// 		data.PenyekebanGalungan,
// 		data.PenyajanGalungan,
// 		data.PenampahanGalungan,
// 		data.Galungan,
// 		data.ManisGalungan,
// 		data.PemaridanGuru,
// 		data.Ulihan,
// 		data.PemacekanAgung,
// 		data.PenampahanKuningan,
// 		data.Kuningan,
// 		data.PegatUwakan,
// 		data.TumpekKandang,
// 		data.TumpekWayang,
// 		data.Saraswati,
// 		data.BanyuPinaruh,
// 		data.SiwaRatri,
// 		data.TawurAgungKasanga,
// 		data.Nyepi,
// 		data.NgembakGeni,
// 		data.BudaCemeng,
// 		data.AnggaraKasih,
// 		data.KajengKliwon,
// 		data.Purnama,
// 		data.Tilem,
// 	}
// }

/*
Get Rahinan for Saka instance
*/
func (s *Saka) Rahinan() []Data {
	if s.rahinan == nil {
		arr := []Data{}

		if s.TriWara == Enum.TriWara.Kajeng && s.PancaWara == Enum.PancaWara.Kliwon {
			arr = append(arr, Enum.Rahinan.KajengKliwon)
		}

		if s.SaptaWara == Enum.SaptaWara.Anggara && s.PancaWara == Enum.PancaWara.Kliwon {
			arr = append(arr, Enum.Rahinan.AnggaraKasih)
		} else if s.SaptaWara == Enum.SaptaWara.Buda && s.PancaWara == Enum.PancaWara.Wage {
			arr = append(arr, Enum.Rahinan.BudaCemeng)
		}

		if s.Wuku == Enum.Wuku.Sinta {
			if s.SaptaWara == Enum.SaptaWara.Redite {
				arr = append(arr, Enum.Rahinan.BanyuPinaruh)
			} else if s.SaptaWara == Enum.SaptaWara.Soma {
				arr = append(arr, Enum.Rahinan.SomaRibek)
			} else if s.SaptaWara == Enum.SaptaWara.Anggara {
				arr = append(arr, Enum.Rahinan.SabuhEmas)
			} else if s.SaptaWara == Enum.SaptaWara.Buda {
				arr = append(arr, Enum.Rahinan.PagerWesi)
			}
		} else if s.Wuku == Enum.Wuku.Landep {
			if s.SaptaWara == Enum.SaptaWara.Saniscara {
				arr = append(arr, Enum.Rahinan.TumpekLandep)
			}
		} else if s.Wuku == Enum.Wuku.Wariga {
			if s.SaptaWara == Enum.SaptaWara.Saniscara {
				arr = append(arr, Enum.Rahinan.TumpekUduh)
			}
		} else if s.Wuku == Enum.Wuku.Sungsang {
			if s.SaptaWara == Enum.SaptaWara.Wraspati {
				arr = append(arr, Enum.Rahinan.SugihanJawa)
			} else if s.SaptaWara == Enum.SaptaWara.Sukra {
				arr = append(arr, Enum.Rahinan.SugihanBali)
			}
		} else if s.Wuku == Enum.Wuku.Dungulan {
			if s.SaptaWara == Enum.SaptaWara.Redite {
				arr = append(arr, Enum.Rahinan.PenyekebanGalungan)
			} else if s.SaptaWara == Enum.SaptaWara.Soma {
				arr = append(arr, Enum.Rahinan.PenyajanGalungan)
			} else if s.SaptaWara == Enum.SaptaWara.Anggara {
				arr = append(arr, Enum.Rahinan.PenampahanGalungan)
			} else if s.SaptaWara == Enum.SaptaWara.Buda {
				arr = append(arr, Enum.Rahinan.Galungan)
			} else if s.SaptaWara == Enum.SaptaWara.Wraspati {
				arr = append(arr, Enum.Rahinan.ManisGalungan)
			} else if s.SaptaWara == Enum.SaptaWara.Saniscara {
				arr = append(arr, Enum.Rahinan.PemaridanGuru)
			}
		} else if s.Wuku == Enum.Wuku.Kuningan {
			if s.SaptaWara == Enum.SaptaWara.Redite {
				arr = append(arr, Enum.Rahinan.Ulihan)
			} else if s.SaptaWara == Enum.SaptaWara.Soma {
				arr = append(arr, Enum.Rahinan.PemacekanAgung)
			} else if s.SaptaWara == Enum.SaptaWara.Sukra {
				arr = append(arr, Enum.Rahinan.PenampahanKuningan)
			} else if s.SaptaWara == Enum.SaptaWara.Saniscara {
				arr = append(arr, Enum.Rahinan.Kuningan)
			}
		} else if s.Wuku == Enum.Wuku.Pahang {
			if s.SaptaWara == Enum.SaptaWara.Buda {
				arr = append(arr, Enum.Rahinan.PegatUwakan)
			}
		} else if s.Wuku == Enum.Wuku.Uye {
			if s.SaptaWara == Enum.SaptaWara.Saniscara {
				arr = append(arr, Enum.Rahinan.TumpekKandang)
			}
		} else if s.Wuku == Enum.Wuku.Wayang {
			if s.SaptaWara == Enum.SaptaWara.Saniscara {
				arr = append(arr, Enum.Rahinan.TumpekWayang)
			}
		} else if s.Wuku == Enum.Wuku.Watugunung {
			if s.SaptaWara == Enum.SaptaWara.Saniscara {
				arr = append(arr, Enum.Rahinan.Saraswati)
			}
		}

		temp := s.Carbon

		n1Day := New(temp.AddDay())
		b1Day := New(temp.SubDay())
		b2Day := New(temp.SubDays(2))

		// fmt.Println(s.Carbon.ToDateString(), b1Day.Saka, s.Saka)
		if n1Day.Sasih == Enum.Sasih.Kapitu && n1Day.SasihDayInfo == Enum.SasihDayInfo.Tilem {
			arr = append(arr, Enum.Rahinan.SiwaRatri)
		} else if s.Saka < n1Day.Saka {
			arr = append(arr, Enum.Rahinan.TawurAgungKasanga)
		} else if b1Day.Saka < s.Saka {
			arr = append(arr, Enum.Rahinan.Nyepi)
		} else if b2Day.Saka < s.Saka && b1Day.Saka == s.Saka {
			arr = append(arr, Enum.Rahinan.NgembakGeni)
		}

		if s.SasihDayInfo == Enum.SasihDayInfo.Purnama {
			arr = append(arr, Enum.Rahinan.Purnama)
		} else if s.SasihDayInfo == Enum.SasihDayInfo.Tilem {
			arr = append(arr, Enum.Rahinan.Tilem)
		}

		s.rahinan = &arr
	}

	return *s.rahinan
}
