/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-module/carbon"
)

func TestEqual(t *testing.T) {
	saka := New()
	saka2 := New()

	if saka.PancaWara != saka2.PancaWara {
		t.Errorf("Saka and saka2 should be equal %s", saka.PancaWara.Name)
	}
}

func TestInvalidFormat(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Saka must panic")
		}
	}()
	_ = New("Invalid format")
	t.Error("Saka must panic")
}

func TestStringFormat(t *testing.T) {
	calendar := New("2023-03-22")
	nyepi := false
	rangeDay := []Data{}
checkNyepi:
	for _, r := range calendar.Rahinan() {
		rangeDay = append(rangeDay, r)
		if r == Enum.Rahinan.Nyepi {
			nyepi = true
			break checkNyepi
		}
	}

	if !nyepi {
		t.Errorf("This date must nyepi, get %+v", rangeDay)
	}
}

func TestBytesFormat(t *testing.T) {
	calendar := New([]byte("2023-03-22"))
	nyepi := false
	rangeDay := []Data{}
checkNyepi:
	for _, r := range calendar.Rahinan() {
		rangeDay = append(rangeDay, r)
		if r == Enum.Rahinan.Nyepi {
			nyepi = true
			break checkNyepi
		}
	}

	if !nyepi {
		t.Errorf("This date must nyepi, get %+v", rangeDay)
	}
}

func TestTimeFormat(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Error(err)
	}
	calendar := New(time.Date(2023, 3, 22, 0, 0, 0, 0, loc))
	nyepi := false
	rangeDay := []Data{}
checkNyepi:
	for _, r := range calendar.Rahinan() {
		rangeDay = append(rangeDay, r)
		if r == Enum.Rahinan.Nyepi {
			nyepi = true
			break checkNyepi
		}
	}

	if !nyepi {
		t.Errorf("This date must nyepi, get %+v", rangeDay)
	}
}

func TestNyepi(t *testing.T) {
	calendar := New(carbon.CreateFromDate(2023, 3, 22))
	nyepi := false

checkNyepi:
	for _, r := range calendar.Rahinan() {
		if r == Enum.Rahinan.Nyepi {

			nyepi = true
			break checkNyepi
		}
	}

	if !nyepi {
		t.Error("This date must nyepi")
	}
	fmt.Printf("Nyepi saka %d\n", calendar.Saka)
}

func TestSiwaratri(t *testing.T) {
	calendar := New(carbon.CreateFromDate(2023, 1, 20))

	isTrue := false

check:
	for _, r := range calendar.Rahinan() {
		if r == Enum.Rahinan.SiwaRatri {

			isTrue = true
			break check
		}
	}

	if !isTrue {
		t.Error("This date must siwaratri")
	}
}

func TestGalungan(t *testing.T) {
	calendar := New(carbon.CreateFromDate(2023, 8, 2))

	isTrue := false

check:
	for _, r := range calendar.Rahinan() {
		if r == Enum.Rahinan.Galungan {

			isTrue = true
			break check
		}
	}

	if !isTrue {
		t.Error("This date must galungan")
	}
}

func TestKuningan(t *testing.T) {
	calendar := New(carbon.CreateFromDate(2023, 8, 12))

	isTrue := false

check:
	for _, r := range calendar.Rahinan() {
		if r == Enum.Rahinan.Kuningan {

			isTrue = true
			break check
		}
	}

	if !isTrue {
		t.Error("This date must kuningan")
	}
}
