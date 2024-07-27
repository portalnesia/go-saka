/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type EkaJalaRsi struct {
	BagnaMapasah       Data
	BahuPutra          Data
	BahuAstawa         Data
	BuatKingking       Data
	BuatLara           Data
	BuatMerang         Data
	BuatSebet          Data
	BuatSuka           Data
	DahatKingking      Data
	Kameranan          Data
	Kamertaan          Data
	Kasobagian         Data
	KinasihanAmerta    Data
	KinasihanJana      Data
	LanggengKayohanaan Data
	LuwihBagia         Data
	ManggihBagia       Data
	ManggihSuka        Data
	PatiningAmerta     Data
	Rahayu             Data
	SidhaKasobagian    Data
	Subagia            Data
	SukaKapanggih      Data
	SukaPinanggih      Data
	SukaRahayu         Data
	TininggalingSuka   Data
	WerdhiPutra        Data
	WerdhiSarwaMule    Data
}

func newEkaJalaRsi() EkaJalaRsi {
	return EkaJalaRsi{
		BagnaMapasah:       Data{0, "Bagna Mapasah"},
		BahuPutra:          Data{1, "Bahu Putra"},
		BahuAstawa:         Data{2, "Bahu Astawa"},
		BuatKingking:       Data{3, "Buat Kingking"},
		BuatLara:           Data{4, "Buat Lara"},
		BuatMerang:         Data{5, "Buat Merang"},
		BuatSebet:          Data{6, "Buat Sebet"},
		BuatSuka:           Data{7, "Buat Suka"},
		DahatKingking:      Data{8, "Dahat Kingking"},
		Kameranan:          Data{9, "Kameranan"},
		Kamertaan:          Data{10, "Kamertaan"},
		Kasobagian:         Data{11, "Kasobagian"},
		KinasihanAmerta:    Data{12, "Kinasihan Amerta"},
		KinasihanJana:      Data{13, "Kinasihan Jana"},
		LanggengKayohanaan: Data{14, "Langgeng Kayohanaan"},
		LuwihBagia:         Data{15, "Luwih Bagia"},
		ManggihBagia:       Data{16, "Manggih Bagia"},
		ManggihSuka:        Data{17, "Manggih Suka"},
		PatiningAmerta:     Data{18, "Patining Amerta"},
		Rahayu:             Data{19, "Rahayu"},
		SidhaKasobagian:    Data{20, "Sidha Kasobagian"},
		Subagia:            Data{21, "Subagia"},
		SukaKapanggih:      Data{22, "Suka Kapanggih"},
		SukaPinanggih:      Data{23, "Suka Pinanggih"},
		SukaRahayu:         Data{24, "Suka Rahayu"},
		TininggalingSuka:   Data{25, "Tininggaling Suka"},
		WerdhiPutra:        Data{26, "Werdhi Putra"},
		WerdhiSarwaMule:    Data{27, "Werdhi Sarwa Mule"},
	}
}

func newEkaJalaRsiSlice() []Data {
	data := newEkaJalaRsi()
	return []Data{
		data.BagnaMapasah,
		data.BahuPutra,
		data.BahuAstawa,
		data.BuatKingking,
		data.BuatLara,
		data.BuatMerang,
		data.BuatSebet,
		data.BuatSuka,
		data.DahatKingking,
		data.Kameranan,
		data.Kamertaan,
		data.Kasobagian,
		data.KinasihanAmerta,
		data.KinasihanJana,
		data.LanggengKayohanaan,
		data.LuwihBagia,
		data.ManggihBagia,
		data.ManggihSuka,
		data.PatiningAmerta,
		data.Rahayu,
		data.SidhaKasobagian,
		data.Subagia,
		data.SukaKapanggih,
		data.SukaPinanggih,
		data.SukaRahayu,
		data.TininggalingSuka,
		data.WerdhiPutra,
		data.WerdhiSarwaMule,
	}
}
