package currency

import (
	"menu/common/database"
	"menu/common/model"
)

func ConvertPrice(from model.Currency, to model.Currency, price *model.Price) {
	fromStock := from.Exchange_rate
	toStock := to.Exchange_rate

	pBuff := price.GetFloat64()

	pBuff /= fromStock
	pBuff *= toStock

	*price = model.NewPriceFromFloat(pBuff)
}

func GetMainCurrency() *model.Currency {
	c := &model.Currency{}
	database.DB.Where("exchange_rate = ?", 1.0000).First(c)
	return c
}
