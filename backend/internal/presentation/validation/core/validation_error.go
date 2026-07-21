package core

type ValidationError struct {
	field string
	tag   string
	param string
}

func NewValidationError(
	field string,
	tag string,
	param string,
) ValidationError {

	return ValidationError{
		field: field,
		tag:   tag,
		param: param,
	}
}

func (v ValidationError) Error() string {

	if v.param != "" {
		return v.field + " failed on the " + v.tag + " validation with parameter " + v.param
	}

	return v.field + " failed on the " + v.tag + " validation"
}

func (v ValidationError) Field() string {
	return v.field
}

func (v ValidationError) Tag() string {
	return v.tag
}

func (v ValidationError) Param() string {
	return v.param
}
