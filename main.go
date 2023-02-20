package main

import (
	"flag"
	"fmt"
)

var fahr float64
var kel float64
var cel float64
var out string
var funfacts string

func init() {

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")

}

func main() {

	flag.Parse()

	// kelvin to celsius
	if isFlagPassed("K") && out == "C" {
		output := cel + 273.15
		fmt.Printf("%v K er %v C", kel, output)
	}
	// cel to kel
	if isFlagPassed("C") && out == "K" {
		output := cel + 273.15
		fmt.Printf("%v C er %v K", cel, output)
	}

	// cel to far
	if isFlagPassed("C") && out == "F" {
		output := cel*(9/5) + 32
		fmt.Printf("%v C er %v F", cel, output)

	}
	// far to cel
	if isFlagPassed("F") && out == "C" {
		output := (fahr - 32) * 5 / 9
		fmt.Printf("%v F er %v C", fahr, output)
	}

	// kel to far
	if isFlagPassed("F") && out == "C" {
		output := (kel-273.15)*(9/5) + 32
		fmt.Printf("%v F er %v C", fahr, output)
	}
	// far to kel
	if isFlagPassed("K") && out == "F" {
		output := (fahr-32)*(5/9) + 273.15
		fmt.Printf("%v K er %v F", kel, output)
	}

	if isFlagPassed("funfacts") {
		funfactsAbout := GetFunFacts(funfacts)
		for i := 0; i < len(funfactsAbout); i++ {
			fmt.Println(funfactsAbout[i])
		}
	}

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {
	funFacts := new(FunFacts)

	funFacts.Sun = []string{"Temperatur i Solens kjerne 15000000 C", "Temperatur på ytre lag av Solen 5778 K"}
	funFacts.Luna = []string{"Temperatur på Månens overflate om natten -183 C", "Temperatur på Månens overflate om dagen 106 C"}
	funFacts.Terra = []string{"Høyeste temperatur målt på Jordens overflate er 56.7 C", "Laveste temperatur målt på Jordens overflate -89.4 C", "Temperatur i Jordens indre kjerne 9392 K"}

	if about == "sun" {
		return funFacts.Sun
	} else if about == "luna" {
		return funFacts.Luna
	} else {
		return funFacts.Terra
	}
}
