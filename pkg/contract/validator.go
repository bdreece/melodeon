package contract

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Validator struct {
    validator validator.Validate
}

func (v *Validator) Validate(i any) error {
    if err := v.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    return nil
}

func NewValidator() *Validator {
    v := validator.New()
    return &Validator{*v}
}

var _ echo.Validator = (*Validator)(nil)
