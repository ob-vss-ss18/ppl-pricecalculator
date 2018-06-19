package main

import (
	"net/http"
	"fmt"
)

/*
	Calculation of the total price.
 */
func calculateTotalPrice(items []Item) (float64, error) {

	//TODO calculate the total price

	totalPrice := 0.0

	//feste Werte
	maxGebraucht := 149.0
	maxNeu := 249.0
	//testWert := 50.0 // Wert von Testen von anderen Skis
	serviceFlat := 0.0 //49.0
	versicherung := 0.0 //30.0

	for i := range items {

		//eingaben
		neuwert := items[i].PriceNew
		amortisationFactor := items[i].Amortisation_factor
		rabattFamilie := items[i].Family_discount
		rabattSaison := items[i].Discout_perc
		gebraucht := items[i].Condition == 1
		zusatz := items[i].Additional_stuff //stoecke etc.

		zwischenErgebnis := neuwert * amortisationFactor

		if gebraucht && zwischenErgebnis > maxGebraucht {
			zwischenErgebnis = maxGebraucht
		}
		if !gebraucht && zwischenErgebnis > maxNeu {
			zwischenErgebnis = maxNeu
		}

		totalPrice += (zwischenErgebnis + serviceFlat + versicherung +
			zusatz) * (1.0 - rabattFamilie) * (1.0 - rabattSaison)
		//reference := neuwert + testWert + serviceFlat + versicherung + zusatz

	}

	return totalPrice, nil
}

/*
	calculation price with static values. Display price in browser.
 */
func calculateExamplePrice(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "hello, world")
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
