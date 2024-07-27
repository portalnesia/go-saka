/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type TriWara struct {
	Pasah  Wara
	Beteng Wara
	Kajeng Wara
}

func newTriWara() TriWara {
	return TriWara{
		Pasah:  Wara{0, 9, "Pasag"},
		Beteng: Wara{1, 4, "Beteng"},
		Kajeng: Wara{2, 7, "Kajeng"},
	}
}

func newTwiWaraSlice() []Wara {
	data := newTriWara()
	return []Wara{data.Pasah, data.Beteng, data.Kajeng}
}
