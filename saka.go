/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

import (
	"math"
	"time"

	"github.com/dromara/carbon/v2"
)

type enumData struct {
	AstaWara          AstaWara
	CaturWara         CaturWara
	DasaWara          DasaWara
	DwiWara           DwiWara
	EkaJalaRsi        EkaJalaRsi
	EkaWara           EkaWara
	Ingkel            Ingkel
	Jejapan           Jejapan
	Lintang           Lintang
	PancaWara         PancaWara
	PancaSuda         PancaSuda
	Pararasan         Pararasan
	Pivot             Pivot
	PratithiSamutPada PratithiSamutPada
	Rahinan           Rahinan
	Rakam             Rakam
	SadWara           SadWara
	SangaWara         SangaWara
	SaptaWara         SaptaWara
	Sasih             Sasih
	SasihDayInfo      SasihDayInfo
	TriWara           TriWara
	WatekAlit         WatekAlit
	WatekMadya        WatekMadya
	Wuku              Wuku
}

var Enum = enumData{
	AstaWara:          newAstaWara(),
	CaturWara:         newCaturWara(),
	DasaWara:          newDasaWara(),
	DwiWara:           newDwiWara(),
	EkaJalaRsi:        newEkaJalaRsi(),
	EkaWara:           newEkaWara(),
	Ingkel:            newIngkel(),
	Jejapan:           newJejapan(),
	Lintang:           newLintang(),
	PancaWara:         newPancaWara(),
	PancaSuda:         newPancaSuda(),
	Pararasan:         newPararasan(),
	Pivot:             newPivot(),
	PratithiSamutPada: newPratithiSamutPada(),
	Rahinan:           newRahinan(),
	Rakam:             newRakam(),
	SadWara:           newSadWara(),
	SangaWara:         newSangaWara(),
	SaptaWara:         newSaptaWara(),
	Sasih:             newSasih(),
	SasihDayInfo:      newSasihDayInfo(),
	TriWara:           newTriWara(),
	WatekAlit:         newWatekAlit(),
	WatekMadya:        newWatekMadya(),
	Wuku:              newWuku(),
}

type Saka struct {
	carbon carbon.Carbon

	Wuku              WukuData
	TriWara           Wara
	PancaWara         PancaWaraData
	SadWara           Wara
	SaptaWara         SaptaWaraData
	Jejapan           Data
	Lintang           Data
	PancaSuda         Data
	Rakam             Data
	EkaJalaRsi        Data
	Ingkel            Data
	EkaWara           Wara
	DwiWara           Wara
	DasaWara          Wara
	WatekAlit         Data
	WatekMadya        Data
	Pararasan         Data
	CaturWara         Wara
	AstaWara          Wara
	SangaWara         Wara
	Saka              int64
	Sasih             SasihData
	SasihDayInfo      SasihDayInfoData
	SasihDay          []int64 // 2 or one array of int64
	PratithiSamutPada Data

	rahinan *[]Data
}

// Carbon get carbon instance for current Saka
func (s *Saka) Carbon() carbon.Carbon {
	return s.carbon
}

// Time get time instance for current Saka
func (s *Saka) Time() time.Time {
	return s.carbon.StdTime()
}

// New Create new Saka instance
//
// Accept any format of date. For example
//
// - 2023-08-03 // valid date string
//
// - time.Now() // time package
//
// - carbon.Now() // carbon package
//
// Default: now
func New(dateArg ...interface{}) *Saka {
	saka := Saka{}

	date := getCarbon(dateArg...)
	saka.carbon = date

	pivot := getBestPivot(date)
	pawukonDay := getPawukonDay(pivot, date)

	saka.Wuku = newWukuSlice()[int64(math.Floor(float64(pawukonDay)/7))]
	saka.TriWara = newTwiWaraSlice()[pawukonDay%3]
	saka.PancaWara = newPancaWaraSlice()[pawukonDay%5]
	saka.SadWara = newSadWaraSlice()[pawukonDay%6]
	saka.SaptaWara = newSaptaWaraSlice()[pawukonDay%7]
	saka.Jejapan = newJejapanSlice()[pawukonDay%6]
	saka.Lintang = newLintangSlice()[pawukonDay%35]
	saka.PancaSuda = newPancaSudaSlice()[(saka.SaptaWara.KertaAji+saka.PancaWara.Urip)%7]
	saka.Rakam = newRakamSlice()[(saka.SaptaWara.Kupih+saka.PancaWara.Kupih)%6]
	saka.EkaJalaRsi = newEkaJalaRsiSlice()[c_ekajalarsislice[pawukonDay]]
	saka.Ingkel = newIngkelSlice()[saka.Wuku.ID%6]

	urip := saka.PancaWara.Urip + saka.SaptaWara.Urip
	saka.EkaWara = newEkaWaraSlice()[urip%2]
	saka.DwiWara = newDwiWaraSlice()[urip%2]
	saka.DasaWara = newDasaWaraSlice()[urip%10]
	saka.WatekAlit = newWatekAlitSlice()[urip%4]
	saka.WatekMadya = newWatekMadyaSlice()[urip%5]
	saka.Pararasan = newPararasanSlice()[urip%10]
	saka.CaturWara = getCaturwara(pawukonDay)
	saka.AstaWara = getAstawara(pawukonDay)
	saka.SangaWara = getSangawara(pawukonDay)

	resSasih := getSasihInfo(pivot, date)
	resSasihDay := getSasihDay(pivot, date)

	saka.Saka = resSasih[0]
	saka.Sasih = getSasih(resSasih)
	saka.SasihDayInfo = getSasihDayInfo(resSasihDay, saka.Sasih, saka.Saka)

	var sasihDay []int64
	if resSasihDay[2] == 1 {
		sasihDay = make([]int64, 2)
		sasihDay[0] = resSasihDay[0]
		if resSasihDay[0] == 15 {
			sasihDay[1] = 1
		} else {
			sasihDay[1] = resSasihDay[0] + 1
		}
	} else {
		sasihDay = make([]int64, 1)
		sasihDay[0] = resSasihDay[0]
	}
	saka.SasihDay = sasihDay

	saka.PratithiSamutPada = getPratithiSamutPada(saka.SasihDay, saka.SasihDayInfo, saka.Sasih, date)

	return &saka
}
