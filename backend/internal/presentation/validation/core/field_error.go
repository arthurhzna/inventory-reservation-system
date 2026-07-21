package core

type FieldError interface {
	error

	Field() string
	Tag() string
	Param() string
}
