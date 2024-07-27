/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type WatekMadya struct {
	Wong  Data
	Gajah Data
	Watu  Data
	Buta  Data
	Suku  Data
}

func newWatekMadya() WatekMadya {
	return WatekMadya{
		Wong:  Data{0, "Wong"},
		Gajah: Data{1, "Gajah"},
		Watu:  Data{2, "Watu"},
		Buta:  Data{3, "Buta"},
		Suku:  Data{4, "Suku"},
	}
}

func newWatekMadyaSlice() []Data {
	data := newWatekMadya()
	return []Data{data.Wong, data.Gajah, data.Watu, data.Buta, data.Suku}
}
