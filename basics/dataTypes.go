package basics

import (
	"fmt"
	"math/cmplx"
)

// Types of data in Golang
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

func DataTypes() {
	var (
		ToBe       bool       = false
		Text       string     = "Hello World"
		Int        int        = 42 // int is a number at least 32 bits in size but not a alias for int32
		Int8       int8       = 127
		Int16      int16      = 32767
		Int32      int32      = 2147483647
		Int64      int64      = 9223372036854775807
		Uint       uint       = 42 // uint is an unsigned integer at least 32 bits in size but not a alias for uint32
		Uint8      uint8      = 255
		Uint16     uint16     = 65535
		Uint32     uint32     = 4294967295
		Uint64     uint64     = 18446744073709551615
		Byte       byte       = 'a' // alias for uint8 used to save char
		Rune       rune       = '€' // alias for int32 used to save Unicode code point
		Float32    float32    = 3.14
		Float64    float64    = 3.14
		Complex64  complex64  = complex64(cmplx.Sqrt(-5 + 12i))
		Complex128 complex128 = cmplx.Sqrt(-5 + 12i)
		Uintptr    uintptr    = 42 //types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", Text, Text)
	fmt.Printf("Type: %T Value: %v\n", Int, Int)
	fmt.Printf("Type: %T Value: %v\n", Int8, Int8)
	fmt.Printf("Type: %T Value: %v\n", Int16, Int16)
	fmt.Printf("Type: %T Value: %v\n", Int32, Int32)
	fmt.Printf("Type: %T Value: %v\n", Int64, Int64)
	fmt.Printf("Type: %T Value: %v\n", Uint, Uint)
	fmt.Printf("Type: %T Value: %v\n", Uint8, Uint8)
	fmt.Printf("Type: %T Value: %v\n", Uint16, Uint16)
	fmt.Printf("Type: %T Value: %v\n", Uint32, Uint32)
	fmt.Printf("Type: %T Value: %v\n", Uint64, Uint64)
	fmt.Printf("Type: %T Value: %c\n", Byte, Byte)
	fmt.Printf("Type: %T Value: %U\n", Rune, Rune)
	fmt.Printf("Type: %T Value: %v\n", Float32, Float32)
	fmt.Printf("Type: %T Value: %v\n", Float64, Float64)
	fmt.Printf("Type: %T Value: %v\n", Complex64, Complex64)
	fmt.Printf("Type: %T Value: %v\n", Complex128, Complex128)
	fmt.Printf("Type: %T Value: %v\n", Uintptr, Uintptr)
	var ZeroInt int
	var ZeroUint uint
	var ZeroString string
	var ZeroBool bool
	var ZeroFloat float64
	fmt.Printf("Type: %T Value: %v\n", ZeroInt, ZeroInt)
	fmt.Printf("Type: %T Value: %v\n", ZeroUint, ZeroUint)
	fmt.Printf("Type: %T Value: %v\n", ZeroString, ZeroString)
	fmt.Printf("Type: %T Value: %v\n", ZeroBool, ZeroBool)
	fmt.Printf("Type: %T Value: %v\n", ZeroFloat, ZeroFloat)
}
