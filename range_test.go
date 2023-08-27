/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

import (
	"testing"

	"github.com/golang-module/carbon"
)

func TestDateNil(t *testing.T) {
	d1 := carbon.CreateFromDate(2023, 8, 27)

	_, err := NewRange(nil, d1)
	if err == nil {
		t.Error("Error not shown")
	} else if err.Error() != "date1 and date2 cannot be nil" {
		t.Error("Error message invalid")
	}
}

func TestReverseDateRange(t *testing.T) {
	d1 := carbon.CreateFromDate(2023, 8, 27)
	d2 := carbon.CreateFromDate(2023, 8, 31)

	saka, err := NewRange(d2, d1)
	if err != nil {
		t.Error(err)
	}

	rahinan := saka.ListAllRahinan()
	if len(rahinan) != 2 {
		t.Error("Rahinan failed")
	}
}

func TestRangeListAllRahinan(t *testing.T) {
	d1 := carbon.CreateFromDate(2023, 8, 27)
	d2 := carbon.CreateFromDate(2023, 8, 31)

	saka, err := NewRange(d1, d2)
	if err != nil {
		t.Error(err)
	}

	rahinan := saka.ListAllRahinan()
	if len(rahinan) != 2 {
		t.Error("Rahinan failed")
	}

	saka, err = NewRange(d2, d1)
	if err != nil {
		t.Error(err)
	}

	rahinan = saka.ListAllRahinan()
	if len(rahinan) != 2 {
		t.Error("Rahinan failed")
	}
}

func TestRangeListByDate(t *testing.T) {
	d1 := carbon.CreateFromDate(2023, 8, 27)
	d2 := carbon.CreateFromDate(2023, 8, 31)

	saka, err := NewRange(d1, d2)
	if err != nil {
		t.Error(err)
	}

	rahinan := saka.ListByDate()
	for i, r := range rahinan {
		if i == 0 {
			if r.Saka.Carbon.DayOfMonth() != d1.DayOfMonth() {
				t.Errorf("First date not same: %s", r.Saka.Carbon.ToString())
			}
			if r.Rahinan[0].ID != 28 {
				t.Errorf("First rahinan not same: %s", r.Rahinan[0].Name)
			}
		}
		if i == len(rahinan)-1 {
			if r.Saka.Carbon.DayOfMonth() != d2.DayOfMonth() {
				t.Errorf("Last date not same: %s", r.Saka.Carbon.ToString())
			}
			if r.Rahinan[0].ID != 29 {
				t.Errorf("Last rahinan not same: %s", r.Rahinan[0].Name)
			}
		}
	}
}

func TestMaxRangeError(t *testing.T) {
	d1 := carbon.CreateFromDate(2023, 8, 5)
	d2 := carbon.CreateFromDate(2023, 8, 31)

	_, err := NewRange(d1, d2)
	if err == nil {
		t.Error("Error not shown")
	} else if err.Error() != "maximum range is 7 days" {
		t.Error("Error message invalid")
	}
}
