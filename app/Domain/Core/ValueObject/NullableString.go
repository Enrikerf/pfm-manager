package ValueObject

type NullableString interface {
	IsNull() bool
	GetValue() string
}

func NewNullableString(isNull bool, value string) NullableString {
	return &nullableString{isNull: isNull, value: value}
}

type nullableString struct {
	isNull bool
	value  string
}

func (n nullableString) IsNull() bool {
	return n.isNull
}

func (n nullableString) GetValue() string {
	return n.value
}
