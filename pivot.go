/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

import "github.com/golang-module/carbon"

type PivotData struct {
	Carbon        carbon.Carbon
	PawukonDay    int64
	SasihDay      int64
	NgunaratriDay int64
	Saka          int64
	Sasih         SasihData
	IsNampihSasih bool
}

type Pivot struct {
	Pivot1971 PivotData
	Pivot2000 PivotData
}

func newPivot() Pivot {

	return Pivot{
		Pivot1971: PivotData{carbon.CreateFromDate(1971, 1, 27), 3, 0, 0, 1892, SasihData{6, 6, "Kapitu", nil}, false},
		Pivot2000: PivotData{carbon.CreateFromDate(2000, 1, 18), 86, 12, 0, 1921, SasihData{6, 6, "Kapitu", nil}, false},
	}
}

// func newPivotSlice() []PivotData {
// 	data := newPivot()
// 	return []PivotData{data.Pivot1971, data.Pivot2000}
// }
