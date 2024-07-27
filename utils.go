/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package saka

import (
	"errors"
	"math"
	"time"

	"github.com/golang-module/carbon"
)

const (
	c_day_pawukon    int64 = 210
	c_day_ngunaratri int64 = 63
)

var pangalantakaPaing = carbon.CreateFromDate(2000, 1, 6)
var skStart = carbon.CreateFromDate(1993, 1, 24)
var skEnd = carbon.CreateFromDate(2003, 1, 3)

var c_ekajalarsislice []int = []int{23, 7, 17, 7, 23, 23, 17, 9, 7, 13, 26, 24, 23, 20, 13, 7, 13, 25, 19, 6, 2,
	14, 26, 17, 20, 25, 22, 0, 10, 6, 15, 23, 7, 17, 23, 17, 25, 5, 23, 2, 2, 2, 12, 12, 5, 14, 12, 26, 26, 1,
	23, 23, 15, 25, 15, 6, 9, 25, 18, 25, 11, 15, 21, 25, 25, 12, 0, 17, 13, 0, 15, 23, 12, 7, 16, 25, 18, 24,
	12, 12, 6, 7, 6, 26, 7, 6, 12, 7, 25, 2, 12, 25, 25, 14, 15, 26, 7, 12, 20, 7, 6, 25, 25, 6, 13, 25, 17, 13,
	23, 6, 26, 20, 25, 25, 23, 7, 18, 18, 17, 7, 17, 7, 5, 26, 17, 6, 9, 12, 12, 13, 25, 18, 18, 6, 2, 25, 25,
	2, 25, 17, 20, 14, 27, 23, 17, 8, 25, 17, 6, 17, 7, 3, 15, 18, 25, 2, 7, 13, 25, 20, 7, 15, 15, 23, 7, 8,
	24, 2, 12, 9, 24, 24, 17, 24, 20, 7, 12, 12, 14, 18, 25, 20, 5, 18, 5, 20, 26, 12, 23, 18, 17, 17, 25, 15,
	2, 24, 4, 2, 23, 25, 18, 25, 20, 14, 4, 2, 25, 7, 25, 17}

func getCarbon(argument ...interface{}) carbon.Carbon {
	var data carbon.Carbon

	if len(argument) == 0 {
		data = carbon.Now()
	} else if len(argument) > 0 {
		switch v := argument[0].(type) {
		case string:
			data = carbon.Parse(v)
		case carbon.Carbon:
			data = v
		case []byte:
			data = carbon.Parse(string(v))
		case time.Time:
			data = carbon.FromStdTime(v)
		default:

		}
	}
	if !data.IsValid() {
		panic(errors.New("invalid date"))
	}
	now := carbon.CreateFromDate(data.Year(), data.Month(), data.DayOfMonth())
	return now
}

func getBestPivot(carbon carbon.Carbon) PivotData {
	if pangalantakaPaing.DiffInDays(carbon) < 0 {
		return Enum.Pivot.Pivot1971
	}
	return Enum.Pivot.Pivot2000
}

func mod(a int64, b int64) int64 {
	return ((a % b) + b) % b
}

func delta(a carbon.Carbon, b carbon.Carbon) int64 {
	return a.DiffInDays(b)
}

func getPawukonDay(pivot PivotData, date carbon.Carbon) int64 {
	return mod(int64(pivot.PawukonDay)+delta(pivot.Carbon, date), c_day_pawukon)
}

func getCaturwara(pawukonDay int64) Wara {
	caturwara := newCaturWaraSlice()
	if pawukonDay < 71 {
		return caturwara[pawukonDay%4]
	} else if pawukonDay > 72 {
		return caturwara[(pawukonDay-2)%4]
	}
	return Enum.CaturWara.Jaya
}

func getAstawara(pawukonDay int64) Wara {
	astawara := newAstaWaraSlice()
	if pawukonDay < 71 {
		return astawara[pawukonDay%8]
	} else if pawukonDay > 72 {
		return astawara[(pawukonDay-2)%8]
	}
	return Enum.AstaWara.Kala
}

func getSangawara(pawukonDay int64) Wara {
	if pawukonDay > 3 {
		return newSangaWaraSlice()[(pawukonDay-3)%9]
	}
	return Enum.SangaWara.Dangu
}

func getSasihInfo(pivot PivotData, date carbon.Carbon) []int64 {
	res := make([]int64, 3)
	pTime := pivot.Carbon
	dayDiff := delta(pTime, date)
	daySkip := int64(math.Ceil(float64(dayDiff) / float64(c_day_ngunaratri)))
	dayTotal := pivot.SasihDay + dayDiff + daySkip

	pivotOffset := int64(1)
	if pivot.SasihDay == 0 && pivot.NgunaratriDay == 0 {
		pivotOffset = 0
	}

	totalSasih := int64(math.Ceil(float64(dayTotal)/30)) - pivotOffset

	currentSasih := pivot.Sasih.ID
	currentSaka := pivot.Saka
	if currentSasih == Enum.Sasih.Kadasa.ID {
		currentSaka = pivot.Saka - 1
	}
	nampihCount := 0
	if pivot.IsNampihSasih {
		nampihCount = 1
	}
	inSK := pTime.BetweenIncludedStart(skStart, skEnd)
	// inSK := false
	// if skStart.DiffInDays(pTime) >= 0 && skEnd.DiffInDays(pTime) < 0 {
	// 	inSK = true
	// }

	for totalSasih != 0 {
		if dayDiff >= 0 {
			if nampihCount == 0 || nampihCount == 2 {
				nampihCount = 0
				currentSasih = mod(currentSasih+1, 12)
			}
			totalSasih = totalSasih - 1

			if currentSasih == Enum.Sasih.Kadasa.ID && nampihCount == 0 {
				currentSaka = currentSaka + 1
			}

			if currentSasih == Enum.Sasih.Kawolu.ID && currentSaka == 1914 {
				inSK = true
			} else if currentSasih == Enum.Sasih.Kawolu.ID && currentSaka == 1924 {
				inSK = false
			}
		} else {
			if nampihCount == 0 || nampihCount == 2 {
				nampihCount = 0
				currentSasih = mod(currentSasih-1, 12)
			}

			totalSasih = totalSasih + 1

			if currentSasih == Enum.Sasih.Kasanga.ID && nampihCount == 0 {
				currentSaka = currentSaka - 1
			}

			if currentSasih == Enum.Sasih.Kapitu.ID && currentSaka == 1914 {
				inSK = false
			} else if currentSasih == Enum.Sasih.Kapitu.ID && currentSaka == 1924 {
				inSK = true
			}
		}

		switch temp := currentSaka % 19; temp {
		case 0, 6, 11:
			if currentSasih == Enum.Sasih.Destha.ID && !inSK {
				if currentSaka != 1925 {
					nampihCount++
				}
			}
		case 3, 8, 14, 16:
			if currentSasih == Enum.Sasih.Sadha.ID && !inSK {
				nampihCount++
			}
		case 2, 10:
			if currentSasih == Enum.Sasih.Destha.ID && inSK {
				nampihCount++
			}
		case 4:
			if currentSasih == Enum.Sasih.Katiga.ID && inSK {
				nampihCount++
			}
		case 7:
			if currentSasih == Enum.Sasih.Kasa.ID && inSK {
				nampihCount++
			}
		case 13:
			if currentSasih == Enum.Sasih.Kadasa.ID && inSK {
				nampihCount++
			}
		case 15:
			if currentSasih == Enum.Sasih.Karo.ID && inSK {
				nampihCount++
			}
		case 18:
			if currentSasih == Enum.Sasih.Sadha.ID && inSK {
				nampihCount++
			}
		}
	}

	res[0] = currentSaka
	res[1] = currentSasih

	if dayTotal >= 0 {
		if nampihCount == 2 {
			res[2] = 1
		} else {
			res[2] = 0
		}
	} else {
		if nampihCount == 1 {
			res[2] = 1
		} else {
			res[2] = 0
		}
	}

	return res
}

func getSasihDay(pivot PivotData, date carbon.Carbon) []int64 {
	res := make([]int64, 3)

	pTime := pivot.Carbon
	dayDiff := delta(pTime, date)
	daySkip := int64(math.Ceil(float64(dayDiff) / float64(c_day_ngunaratri)))
	dayTotal := pivot.SasihDay + dayDiff + daySkip

	res[0] = mod(int64(dayTotal), 30)
	res[1] = 0
	if res[0] == 0 || res[0] > 15 {
		res[1] = 1
	}
	res[2] = 0
	if mod(dayDiff, c_day_ngunaratri) == 0 {
		res[2] = 1
	}
	res[0] = mod(res[0], 15)
	if res[0] == 0 {
		res[0] = 15
	}

	return res
}

func getSasih(resSasih []int64) SasihData {
	saka := resSasih[0]
	sasih := resSasih[1]
	res := newSasihSlice()[sasih]

	if resSasih[2] == 1 {
		if res == Enum.Sasih.Destha {
			if saka < 1914 {
				res = Enum.Sasih.MalaDestha
			} else {
				res = Enum.Sasih.NampihDestha
			}
		} else if res == Enum.Sasih.Katiga {
			res = Enum.Sasih.NampihKatiga
		} else if res == Enum.Sasih.Kasa {
			res = Enum.Sasih.NampihKasa
		} else if res == Enum.Sasih.Kadasa {
			res = Enum.Sasih.NampihKadasa
		} else if res == Enum.Sasih.Karo {
			res = Enum.Sasih.NampihKaro
		} else if res == Enum.Sasih.Sadha {
			if saka < 1914 {
				res = Enum.Sasih.MalaSadha
			} else {
				res = Enum.Sasih.NampihSadha
			}
		}
	}

	return res
}

func getSasihDayInfo(resSasihDay []int64, sasih SasihData, saka int64) SasihDayInfoData {
	date := resSasihDay[0]
	isPangelong := resSasihDay[1] == 1
	isNgunaratri := resSasihDay[2] == 1

	if isPangelong {
		if date == 15 || (date == 14 && isNgunaratri) {
			return Enum.SasihDayInfo.Tilem
		} else if date == 14 && sasih == Enum.Sasih.Kapitu && saka == 1921 {
			return Enum.SasihDayInfo.Tilem
		} else {
			return Enum.SasihDayInfo.Pangelong
		}
	} else {
		if date == 15 || (date == 14 && isNgunaratri) {
			return Enum.SasihDayInfo.Purnama
		} else {
			return Enum.SasihDayInfo.Penanggal
		}
	}
}

func getPratithiSamutPada(sasihDay []int64, sasihDayInfo SasihDayInfoData, sasih SasihData, date carbon.Carbon) Data {
	move := int64(0)
	isNG := len(sasihDay) > 1
	day := sasihDay[0]
	if isNG {
		day = sasihDay[1]
	}

	if sasihDayInfo.Reference() == Enum.SasihDayInfo.Penanggal {
		if day == 1 && isNG {
			move = 0
		} else {
			if day == 1 && day <= 8 {
				move = day - 1
			} else if day >= 9 && day <= 13 {
				move = day - 2
			} else if day == 14 {
				move = 11
			} else if day == 15 {
				move = 0
			}
		}
	} else {
		if day == 1 && isNG {
			temp := date.AddDay()
			nextDay := New(temp)

			if nextDay.Sasih.Reference() != sasih.Reference() {
				move = -1
			}
		} else {
			if day >= 13 {
				move = day - 11
			} else {
				move = day - 1
			}
		}
	}

	start := newPratithiSamutPadaSlice()[sasih.Reference().ID]
	newID := mod(start.ID-move, 12)

	return newPratithiSamutPadaSlice()[newID]
}
