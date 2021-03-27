package tests

import (
	"menu/common/model"
	"strings"
	"testing"
)

func TestMainCurrency(t *testing.T) {
	c := model.GetMainCurrency()

	if strings.Compare(c.Symbol, "zł") != 0 {
		t.Error("Wrong main currency")
	}
}
