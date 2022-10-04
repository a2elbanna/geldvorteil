package calctax

import "fmt"

var fullPrice float32
var discountedPrice float32

func Getdiscount() (discount float32) {
	fmt.Println("Enter Full Price")
	fmt.Scanln(&fullPrice)
	fmt.Println("Enter discounted price ")
	fmt.Scanln(&discountedPrice)
	discount = fullPrice - discountedPrice
	return discount
}

//Geld Vorteil
func GetzuVersteuernderRabatt() (geldVorteil float32) {
	fullPrice := *&fullPrice
	discountedPrice := *&discountedPrice
	geldVorteil = (fullPrice * (0.96)) - discountedPrice
	return geldVorteil
}

func GetSteueranteil() (steueranteil float32) {
	//TODO switch for more Tax classes, if needed.
	steuersatz := 0.42
	toBeTaxed := GetzuVersteuernderRabatt()
	steueranteil = toBeTaxed * float32(steuersatz)
	return steueranteil
}

//5,5 von Steuer
func GetSoli() (soli float32) {
	steueranteil := GetSteueranteil()
	solisatz := 0.055
	soli = steueranteil * float32(solisatz)
	return soli
}

//3% von geldwerter Vorteil
///2 ??hmm if two persons will Get the benefits
func GetArbeitslosenversicherung() (arbeitslosenversicherung float32) {
	geldwerterVorteil := GetzuVersteuernderRabatt()
	arbeitslosenversicherungAnteil := 0.03
	arbeitslosenversicherung = geldwerterVorteil * float32(arbeitslosenversicherungAnteil)
	return arbeitslosenversicherung
}

//18,7 von geldwerter Vorteil
///2 ??
func GetRentenVersicherung() (rentenVersicherung float32) {
	geldwerterVorteil := GetzuVersteuernderRabatt()
	rentenVersicherungAnteil := 0.187
	rentenVersicherung = geldwerterVorteil * float32(rentenVersicherungAnteil)
	return rentenVersicherung
}

func GetSumSteuerUndSozialver() (sum float32) {
	steuerAnteil := GetSteueranteil()
	soliAnteil := GetSoli()
	arbeitslosenversicherung := GetArbeitslosenversicherung()
	rentenversicherung := GetRentenVersicherung()
	sum = steuerAnteil + soliAnteil + arbeitslosenversicherung + rentenversicherung
	return sum
}

func GetGesamtreisePreis() (gesamtReisePreis float32) {
	umSteuerUndSozialver := GetSumSteuerUndSozialver()
	gesamtReisePreis = *&discountedPrice + umSteuerUndSozialver
	return gesamtReisePreis
}

func GetGesamsersparnis() (gesamtersparnis float32) {
	gesamtersparnis = *&fullPrice - GetGesamtreisePreis()
	return gesamtersparnis
}
