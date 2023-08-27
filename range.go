/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

import (
	"errors"

	"github.com/golang-module/carbon"
)

type SakaRangeDate struct {
	Saka    Saka   // Saka instance
	Rahinan []Data // Array of Rahinan
}

type SakaRange struct {
	data []SakaRangeDate
}

/*
Create new Saka range instance

Accept any format of date. For example

  - 2023-08-03 // valid date string
  - time.Now() // time package
  - carbon.Now() // carbon package
*/
func NewRange(date1 interface{}, date2 interface{}) (*SakaRange, error) {
	if date1 == nil || date2 == nil {
		return nil, errors.New("date1 and date2 cannot be nil")
	}

	d1 := getCarbon(date1)
	d2 := getCarbon(date2)

	rangeDate := d1.DiffAbsInDays(d2) + 1

	if rangeDate > 7 {
		return nil, errors.New("maximum range is 7 days")
	}

	tmpRange := d1.DiffInDays(d2)
	var date carbon.Carbon

	if tmpRange < 0 {
		date = d2
	} else {
		date = d1
	}

	response := []SakaRangeDate{}

	for i := int64(0); i < rangeDate; i++ {
		sk := New(date)
		rahinan := sk.Rahinan()
		response = append(response, SakaRangeDate{Saka: *sk, Rahinan: rahinan})
		date = date.AddDay()
	}

	sakaRange := SakaRange{
		data: response,
	}

	return &sakaRange, nil
}

/*
List all Rahinan in the range by date
*/
func (s *SakaRange) ListByDate() []SakaRangeDate {
	return s.data
}

// List all Rahinan in the range
func (s *SakaRange) ListAllRahinan() []Data {
	var response []Data

	for _, v := range s.data {
		response = append(response, v.Rahinan...)
	}

	return response
}
