package model

import (
	"math"
)

type Price int64

func NewPriceFromFloat(v float64) Price {
	p := Price(math.Round(v * 100))

	return p
}

//multiply only by this function, p = p*n
func (p *Price) Multiply(n Price) *Price {
	pBuff := *p
	pBuff *= n
	*p = Price(math.Round(float64(pBuff) / 100))

	return p
}

func (p *Price) Add(n Price) *Price {
	*p += n

	return p
}

func (p *Price) Subtract(n Price) *Price {
	*p -= n

	return p
}

func (p *Price) Divide(n Price) *Price {
	*p /= Price(math.Round(float64(n) / 100))

	return p
}

func (p *Price) GetFloat64() float64 {
	return float64(*p) / 100.00
}
