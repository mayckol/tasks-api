package validation

type ValidatorInterface interface {
	Validate(data interface{}) ([]InvalidFields, bool)
	ErrorMessage(errs []InvalidFields) map[string]interface{}
}
