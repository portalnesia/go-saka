/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package example

import (
	"fmt"
	"github.com/golang-module/carbon"
	"go.portalnesia.com/saka"
	"time"
)

func RangeDate() {
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
