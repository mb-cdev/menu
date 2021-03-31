package tests

import (
	"fmt"
	"menu/common/currency"
	"menu/common/model"
	"strings"
	"testing"
)

func TestMainCurrency(t *testing.T) {
	c := currency.GetMainCurrency()

	if strings.Compare(c.Symbol, "zł") != 0 {
		t.Error("Wrong main currency")
	}
}

func TestFormattingPrice(t *testing.T) {
	c := model.Currency{
		Symbol:   "$",
		Template: "<symbol><amount>",
	}

	var p model.Price = 12333

	ps := float64(p)
	ps *= 0.3

	fmt.Println(ps)
	fmt.Println(c.FormatPrice(&p))
}

func TestPrice(t *testing.T) {
	p1 := model.NewPriceFromFloat(0.8)
	p2 := model.NewPriceFromFloat(0.08)

	p1.Multiply(p2)

	if p1 != 6 {
		t.Error("Multiply Not valid")
	}
}

func TestPriceConvering(t *testing.T) {
	c0 := model.Currency{
		Name:          "PLN",
		Symbol:        "zł",
		Exchange_rate: 1.0000,
		Template:      "<amount><symbol>",
	}
	c1 := model.Currency{
		Name:          "USD",
		Symbol:        "$",
		Exchange_rate: 0.2520,
		Template:      "<symbol><amount>",
	}
	/*c2 := model.Currency{
		Name:          "EUR",
		Symbol:        "€",
		Exchange_rate: 2145,
		Template:      "<symbol><amount>",
	}*/

	p := model.NewPriceFromFloat(100.00)

	currency.ConvertPrice(c0, c1, &p)
	fmt.Println("Conversion PLN to USD", p)
}
