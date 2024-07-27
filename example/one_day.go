/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package example

import (
	"fmt"
	"go.portalnesia.com/saka"
)

func main() {
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
