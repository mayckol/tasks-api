package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stoewer/go-strcase"
)

type InvalidFields struct {
	Error       bool        `json:"error"`
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
	Reason      []string    `json:"reason"`
}

type Wrapper struct {
	Validator *validator.Validate
}

// NewWrapper returns a new instance of the validation wrapper, with that instance we can create helper functions to validate data.
func NewWrapper() *Wrapper {
	v := validator.New()
	return &Wrapper{
		Validator: v,
	}
}

func (v *Wrapper) Validate(data interface{}) ([]InvalidFields, bool) {
	var validationErrors []InvalidFields
	errs := v.Validator.Struct(data)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			elem := InvalidFields{
				FailedField: strcase.SnakeCase(err.Field()),
				Tag:         err.Tag(),
				Value:       err.Value(),
				Error:       true,
				Reason:      []string{fmt.Sprintf("Field '%s' with value '%v' failed on '%s' validation", err.Field(), err.Value(), err.Tag())},
			}
			validationErrors = append(validationErrors, elem)
		}
		return validationErrors, true
	}

	return validationErrors, false
}

// ErrorMessage returns a map of validation errors.
func (v *Wrapper) ErrorMessage(errs []InvalidFields) map[string]interface{} {
	if len(errs) == 0 {
		return nil
	}

	response := make(map[string]interface{})
	response["message"] = "validation errors occurred"
	response["success"] = false

	invalidFields := make([]map[string]interface{}, 0)

	for _, err := range errs {
		invalidField := map[string]interface{}{
			"field":  strcase.SnakeCase(err.FailedField),
			"reason": err.Reason,
		}
		invalidFields = append(invalidFields, invalidField)
	}

	response["invalid_fields"] = invalidFields

	return response
}
