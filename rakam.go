/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type Rakam struct {
	NujuPati         Data
	KalaTinantang    Data
	DemangKadhuruwan Data
	SanggarWaringin  Data
	MantriSinaroja   Data
	MacanKatawan     Data
}

func newRakam() Rakam {
	return Rakam{
		NujuPati:         Data{0, "Nuju Pati"},
		KalaTinantang:    Data{1, "Kala Tinantang"},
		DemangKadhuruwan: Data{2, "Demang Kadhuruwan"},
		SanggarWaringin:  Data{3, "Sanggar Waringin"},
		MantriSinaroja:   Data{4, "Mantri Sinaroja"},
		MacanKatawan:     Data{5, "Macan Katawan"},
	}
}

func newRakamSlice() []Data {
	data := newRakam()
	return []Data{
		data.NujuPati,
		data.KalaTinantang,
		data.DemangKadhuruwan,
		data.SanggarWaringin,
		data.MantriSinaroja,
		data.MacanKatawan,
	}
}
