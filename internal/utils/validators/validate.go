package validators

type Validator[T any] interface {
	Validate(data T) error
}
