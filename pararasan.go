/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type Pararasan struct {
	LakuPanditaSakti Data
	ArasTuding       Data
	ArasKembang      Data
	LakuBintang      Data
	LakuBulan        Data
	LakuSurya        Data
	LakuAir          Data
	LakuBumi         Data
	LakuApi          Data
	LakuAngin        Data
}

func newPararasan() Pararasan {
	return Pararasan{
		LakuPanditaSakti: Data{0, "Laku Pandita Sakti"},
		ArasTuding:       Data{1, "Aras Tuding"},
		ArasKembang:      Data{2, "Aras Kembang"},
		LakuBintang:      Data{3, "Laku Bintang"},
		LakuBulan:        Data{4, "Laku Bulan"},
		LakuSurya:        Data{5, "Laku Surya"},
		LakuAir:          Data{6, "Laku Air"},
		LakuBumi:         Data{7, "Laku Bumi"},
		LakuApi:          Data{8, "Laku Api"},
		LakuAngin:        Data{9, "Laku Angin"},
	}
}

func newPararasanSlice() []Data {
	data := newPararasan()
	return []Data{
		data.LakuPanditaSakti,
		data.ArasTuding,
		data.ArasKembang,
		data.LakuBintang,
		data.LakuBulan,
		data.LakuSurya,
		data.LakuAir,
		data.LakuBumi,
		data.LakuApi,
		data.LakuAngin,
	}
}
