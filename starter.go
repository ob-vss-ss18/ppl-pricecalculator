package main
//Found in https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html
import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"strconv"
	"math"
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
	maxGebraucht := 149.0;
	maxNeu := 249.0;
	preis := 139.0
	testWert := 50.0;
	serviceFlat := 49.0;
	versicherung := 30.0;
	factor := 0.36;
	rabattFamilie := 0.10;
	rabattSaison := 0.10;
	gebraucht := false
	zusatz := 0.0 //stöcke oder so

	zwischenErgebnis := 100.0;
	result := zwischenErgebnis + (testWert +serviceFlat + versicherung  +
		(preis* factor) + zusatz ) * (1.0-rabattFamilie) * (1.0-rabattSaison)

	if (gebraucht && result > maxGebraucht){
		result = maxGebraucht
	}
	if (!gebraucht && result > maxNeu){
		result = maxNeu
	}

	fmt.Println(res,"Result = %d",result);


}
