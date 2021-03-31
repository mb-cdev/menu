package model

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model
	Name          string
	Symbol        string
	Template      string
	Exchange_rate float64 `gorm:"type:numeric(5,4) not null"`
}

func (c *Currency) FormatPrice(price *Price) string {
	tpl := c.Template
	p := float64(*price) / 100
	pString := strconv.FormatFloat(p, 'f', 2, 64)

	tpl = strings.ReplaceAll(tpl, "<symbol>", c.Symbol)
	tpl = strings.ReplaceAll(tpl, "<amount>", pString)

	return tpl
}
