package currency

import "fmt"

type Currency struct {
    entity string
    name   string
    code   string
    minor_unit int
}

var argCurr = &Currency{"ARGENTINA", "Argentine Peso", "ARS", 2}
var brlCurr = &Currency{"BRAZIL", "Brazilian Real", "BRL", 2}
var ilsCurr = &Currency{"ISRAEL", "New Israeli Sheqel", "ILS", 2}
var usdCurr = &Currency{"UNITED STATES", "US Dollar", "USD", 2}
var currencyList = map[string]*Currency{"ARGENTINA": argCurr, "BRAZIL": brlCurr, "ISRAEL": ilsCurr, "UNITED STATES": usdCurr}

func GetCurrency(countryName string) (curr Currency, err error) {
    cur, ok := currencyList[countryName]
    if !ok {
	err = fmt.Errorf("country %s not found", countryName)
	return
    }
    curr = *cur
    return
}
