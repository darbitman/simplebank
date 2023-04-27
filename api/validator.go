package api

import (
	"github.com/darbitman/simplebank/util"
	v "github.com/go-playground/validator/v10"
)

var validCurrency v.Func = func(fieldLevel v.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
