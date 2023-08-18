package nullable

type nullable interface {
	bool | int | string
}

type Value[T nullable] struct {
	value T
	isSet bool
}

func Null[T nullable]() Value[T] {
	var null T
	return Value[T]{null, false}
}

func From[T nullable](value T) Value[T] {
	return Value[T]{value, true}
}

func FromPtr[T nullable](value *T) Value[T] {
	if value == nil {
		return Null[T]()
	}
	return From[T](*value)
}
