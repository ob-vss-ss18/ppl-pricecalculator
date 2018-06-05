package main
//Found in https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html
import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {



	fmt.Fprintln(res, "hello, world - DEV")
	// WERTE abfrage über Item id
	//
	//feste Werte
	maxGebraucht := 149.0
	maxNeu := 249.0
	testWert := 50.0
	serviceFlat := 49.0
	versicherung := 30.0

	//eingaben
	neuwert := 439.0
	amortisationFactor := 0.36
	rabattFamilie := 0.10
	rabattSaison := 0.10
	gebraucht := false
	zusatz := 0.0 //stoecke oder so

	zwischenErgebnis := (neuwert * amortisationFactor)

	if (gebraucht && zwischenErgebnis > maxGebraucht){
		zwischenErgebnis = maxGebraucht
	}
	if (!gebraucht && zwischenErgebnis > maxNeu){
		zwischenErgebnis = maxNeu
	}

	result := (zwischenErgebnis + serviceFlat + versicherung +
		 zusatz) * (1.0 - rabattFamilie) * (1.0 - rabattSaison)
	reference := neuwert + testWert + serviceFlat + versicherung + zusatz

	fmt.Fprintln(res,"Preis = ",result)

	fmt.Fprintln(res,"Referenz = ", reference)


}
