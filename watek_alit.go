/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type WatekAlit struct {
	Lintah Data
	Uler   Data
	Gajah  Data
	Lembu  Data
}

func newWatekAlit() WatekAlit {
	return WatekAlit{
		Lintah: Data{0, "Lintah"},
		Uler:   Data{1, "Uler"},
		Gajah:  Data{2, "Gajah"},
		Lembu:  Data{3, "Lembu"},
	}
}

func newWatekAlitSlice() []Data {
	data := newWatekAlit()
	return []Data{data.Lintah, data.Uler, data.Gajah, data.Lembu}
}
