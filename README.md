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
    "go.portalnesia.com/saka"
    "github.com/golang-module/carbon"
)

func oneDate() {
    saka_instance := saka.New()
    
    // Get information about the date
    saka_instance.EkaWara
    saka_instance.PancaWara
    saka_instance.Sasih
    // and others...

    all_rahinan := saka_instance.Rahinan()
    // todo with list rahinan
    if all_rahinan[0] == saka.Enum.Rahinan.Nyepi {

    }
    // and others...
    
}

func rangeDate() {
    d1 := carbon.CreateFromDate(2023, 8, 27)
	d2 := carbon.CreateFromDate(2023, 8, 31)

    saka_instance := saka.NewRange(d1,d2)

    // List all rahinan
    all_rahinan := saka_instance.ListAllRahinan()

    // List by date
    date_list := saka_instance.ListByDate()
}
```

## Go Reference

[pkg.go.dev/go.portalnesia.com/saka](https://pkg.go.dev/go.portalnesia.com/saka)

## Thanks To

- [github.com/peradnya/balinese-date-js-lib](https://github.com/peradnya/balinese-date-js-lib) - Balinese calendar library for Javascript/Typescript