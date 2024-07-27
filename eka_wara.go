/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

type EkaWara struct {
	Void  Wara
	Luang Wara
}

func newEkaWara() EkaWara {
	return EkaWara{
		Void:  Wara{0, 0, ""},
		Luang: Wara{1, 1, "Luang"},
	}
}

func newEkaWaraSlice() []Wara {
	data := newEkaWara()
	return []Wara{data.Void, data.Luang}
}
