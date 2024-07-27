/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type AstaWara struct {
	Sri    Wara
	Indra  Wara
	Guru   Wara
	Yama   Wara
	Ludra  Wara
	Brahma Wara
	Kala   Wara
	Uma    Wara
}

func newAstaWara() AstaWara {
	return AstaWara{
		Sri:    Wara{ID: 0, Urip: 6, Name: "Sri"},
		Indra:  Wara{ID: 1, Urip: 5, Name: "Indra"},
		Guru:   Wara{ID: 2, Urip: 8, Name: "Guru"},
		Yama:   Wara{ID: 3, Urip: 9, Name: "Yama"},
		Ludra:  Wara{ID: 4, Urip: 3, Name: "Ludra"},
		Brahma: Wara{ID: 5, Urip: 7, Name: "Brahma"},
		Kala:   Wara{ID: 6, Urip: 1, Name: "Kala"},
		Uma:    Wara{ID: 7, Urip: 4, Name: "Uma"},
	}
}

func newAstaWaraSlice() []Wara {
	astawara := newAstaWara()
	return []Wara{
		astawara.Sri,
		astawara.Indra,
		astawara.Guru,
		astawara.Yama,
		astawara.Ludra,
		astawara.Brahma,
		astawara.Kala,
		astawara.Uma,
	}
}
