package main

import (
	"net/http"
	"fmt"
)

/*
	Calculation of the total price.
 */
func calculateTotalPrice(familyDiscount int, items []Item) (int, error) {

	totalPrice := 42

	//TODO calculate the total price

	/*
	println(len(items))
	println(items[0].ItemType)
	println(items[0].Id)
	println(items[0].Condition)
	println(items[0].PriceNew)
	*/

	return totalPrice, nil
}

/*
	calculation price with static values. Display price in browser.
 */
func calculateExamplePrice(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "hello, world - DEV")
	// WERTE abfrage über Item id
	//
	//feste Werte
	maxGebraucht := 149.0
	maxNeu := 249.0
	testWert := 50.0 // Wert von Testen von anderen Skis
	serviceFlat := 49.0
	versicherung := 30.0

	//eingaben
	neuwert := 439.0
	amortisationFactor := 0.36
	rabattFamilie := 0.10
	rabattSaison := 0.10
	gebraucht := false
	zusatz := 0.0 //stoecke etc.

	zwischenErgebnis := neuwert * amortisationFactor

	if gebraucht && zwischenErgebnis > maxGebraucht {
		zwischenErgebnis = maxGebraucht
	}
	if !gebraucht && zwischenErgebnis > maxNeu {
		zwischenErgebnis = maxNeu
	}

	result := (zwischenErgebnis + serviceFlat + versicherung +
		zusatz) * (1.0 - rabattFamilie) * (1.0 - rabattSaison)
	reference := neuwert + testWert + serviceFlat + versicherung + zusatz

	fmt.Fprintln(res, "Preis = ", result)

	fmt.Fprintln(res, "Referenz = ", reference)

}
