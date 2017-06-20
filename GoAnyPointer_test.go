package GoAnyPointer

import "testing"
import "time"

type TestStruct struct {
	String    *string
	Int       *int
	Int8      *int8
	Int16     *int16
	Int32     *int32
	Int64     *int64
	Uint      *uint
	Uint8     *uint8
	Uint16    *uint16
	Uint32    *uint32
	Uint64    *uint64
	Uintptr   *uintptr
	Float32   *float32
	Float64   *float64
	Bool      *bool
	Slice     *[]int
	Array     *[1]int
	Map       *map[string]int
	Interface *interface{}
	Object    *time.Time
}

func TestAll(t *testing.T) {
	test := &TestStruct{
		String:  Sp("test data"),
		Int:     Ip(-1),
		Int8:    I8p(-8),
		Int16:   I16p(-16),
		Int32:   I32p(-32),
		Int64:   I64p(-64),
		Uint:    Up(1),
		Uint8:   U8p(8),
		Uint16:  U16p(16),
		Uint32:  U32p(32),
		Uint64:  U64p(64),
		Uintptr: Uip(0),
		Float32: F32p(3.1415926),
		Float64: F64p(3.14159265358),
		Bool:    Bp(true),
		Slice:   &[]int{1, 2, 3},
		Array:   &[1]int{10},
		Map: &map[string]int{
			"first":  1,
			"second": 2,
		},
		Interface: Inp("interface"),
		Object: (func(o time.Time) *time.Time { return &o })(time.Now()),
	}

	t.Log(*test)
}
