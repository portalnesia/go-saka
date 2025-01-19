[![Go Reference](https://pkg.go.dev/badge/go.portalnesia.com/saka.svg)](https://pkg.go.dev/go.portalnesia.com/saka) ![Go](https://github.com/portalnesia/go-saka/actions/workflows/unit_test.yml/badge.svg)

# Saka

Saka is a Go library for working with the Balinese calendar system. This library provides functionalities to convert Gregorian dates to Balinese calendar dates and vice versa. It also offers features to retrieve Hindu holidays, important cultural events, and auspicious days.

## Installation

To use Saka in your Go project, you need to install it using go get:

```bash
go get -u go.portalnesia.com/saka
```

## Usage

```go
package main

import (
	"fmt"
	"time"
	"github.com/dromara/carbon/v2"
	"go.portalnesia.com/saka"
)

func oneDate() {
	sakaInstance := saka.New()

	// Get information about the date
	fmt.Println("Eka Wara", sakaInstance.EkaWara)
	fmt.Println("Panca Wara", sakaInstance.PancaWara)
	fmt.Println("Sasih", sakaInstance.Sasih)

	// and others...

	allRahinan := sakaInstance.Rahinan()
	for _, rahinan := range allRahinan {
		// todo with list rahinan
		if rahinan == saka.Enum.Rahinan.Nyepi {

		}
	}

}

func rangeDate() {
	d1 := carbon.CreateFromDate(2023, 8, 27)
	d2 := carbon.CreateFromDate(2023, 8, 31)

	sakaInstance, _ := saka.NewRange(d1, d2)

	// List all rahinan in range d1 - d2
	allRahinan := sakaInstance.ListAllRahinan()
	for _, r := range allRahinan {
		fmt.Println(r.Name)
	}

	// List by date
	sakaList := sakaInstance.ListByDate()
	for _, s := range sakaList {
		fmt.Printf("Rahinan for %s\n", s.Time().Format(time.RFC3339))

		for _, r := range s.Rahinan() {
			fmt.Println(r.Name)
		}
		
		fmt.Println()
	}
}
```

## Go Reference

[pkg.go.dev/go.portalnesia.com/saka](https://pkg.go.dev/go.portalnesia.com/saka)

## Thanks To

- [github.com/peradnya/balinese-date-js-lib](https://github.com/peradnya/balinese-date-js-lib) - Balinese calendar library for Javascript/Typescript