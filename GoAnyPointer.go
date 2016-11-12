package GoAnyPointer

func Sp(data string) *string {
	return &data
}

func Ip(data int) *int {
	return &data
}

func I8p(data int8) *int8 {
	return &data
}

func I16p(data int16) *int16 {
	return &data
}

func I32p(data int32) *int32 {
	return &data
}

func I64p(data int64) *int64 {
	return &data
}

func Up(data uint) *uint {
	return &data
}

func U8p(data uint8) *uint8 {
	return &data
}

func U16p(data uint16) *uint16 {
	return &data
}

func U32p(data uint32) *uint32 {
	return &data
}

func U64p(data uint64) *uint64 {
	return &data
}

func Uip(data uintptr) *uintptr {
	return &data
}

func F32p(data float32) *float32 {
	return &data
}

func F64p(data float64) *float64 {
	return &data
}

func Bp(data bool) *bool {
	return &data
}

func Inp(data interface{}) *interface{} {
	return &data
}
