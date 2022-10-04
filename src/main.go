package main

import (
	"fmt"
	"benefitCalculator/pkg"
)

//This programm is to calculate the net profit for companies benefits!

const ()

//var anzahlPerson int

func main() {

	fmt.Println("Discount:", calctax.Getdiscount())
	fmt.Println(" ")
	fmt.Println("####################Result####################")
	fmt.Println("GeldVorteil/benefit :", calctax.GetzuVersteuernderRabatt(), " EURO")
	fmt.Println("Steueranteil für Abrechnung/Tax: ", calctax.GetSteueranteil(), " EURO")
	fmt.Println("solidarität Anteil: ", calctax.GetSoli(), " EURO")
	fmt.Println("Arbeitslosenversicherung: ", calctax.GetArbeitslosenversicherung(), " EURO")
	fmt.Println("RentenVersicherung: ", calctax.GetRentenVersicherung(), " EURO")
	fmt.Println("Sum Steuer Und Sozialver: ", calctax.GetSumSteuerUndSozialver(), " EURO")
	fmt.Println("Actual Full Price: ", calctax.GetGesamtreisePreis(), " EURO")
	fmt.Println("Gesamsersparnis/What you really save: ", calctax.GetGesamsersparnis(), " EURO")
}
