/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type DwiWara struct {
	Menga Wara
	Pepet Wara
}

func newDwiWara() DwiWara {
	return DwiWara{
		Menga: Wara{0, 5, "Menga"},
		Pepet: Wara{1, 4, "Pepet"},
	}
}

func newDwiWaraSlice() []Wara {
	wara := newDwiWara()
	return []Wara{
		wara.Menga,
		wara.Pepet,
	}
}
