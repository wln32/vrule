package ruleimpl

type UnsignedNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type SignedNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Integer interface {
	UnsignedNumber | SignedNumber
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}
