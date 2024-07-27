/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type SadWara struct {
	Tungleh Wara
	Aryang  Wara
	Urukung Wara
	Paniron Wara
	Was     Wara
	Maulu   Wara
}

func newSadWara() SadWara {
	return SadWara{
		Tungleh: Wara{0, 7, "Tungleh"},
		Aryang:  Wara{1, 6, "Aryang"},
		Urukung: Wara{2, 5, "Urukung"},
		Paniron: Wara{3, 8, "Paniron"},
		Was:     Wara{4, 9, "Was"},
		Maulu:   Wara{5, 3, "Maulu"},
	}
}

func newSadWaraSlice() []Wara {
	data := newSadWara()
	return []Wara{data.Tungleh, data.Aryang, data.Urukung, data.Paniron, data.Was, data.Maulu}
}
