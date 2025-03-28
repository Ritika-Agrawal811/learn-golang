package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	/* Boolean */
	var isReady bool = false
	fmt.Println("Boolean:", isReady)

	/* String */
	var name string = "Ritika"
	fmt.Println("String:", name)

	/* Byte (Alias for uint8) */
	var b byte = 'A' // ASCII value of 'A' is 65
	fmt.Println("Byte:", b)

	/* Integer types */
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)

	/* Unsigned Integer types */
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	fmt.Println("uint8:", u8)
	fmt.Println("uint16:", u16)
	fmt.Println("uint32:", u32)
	fmt.Println("uint64:", u64)

	/* Float types */
	var f32 float32 = 3.1415927
	var f64 float64 = 3.141592653589793

	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)

	/* Complex Numbers */
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i

	fmt.Println("complex64:", c64)
	fmt.Println("complex128:", c128)

	/* Rune (Alias for int32, represents Unicode characters) */
	var r rune = '🔥' // Unicode character
	fmt.Println("Rune:", r, string(r))

	/* Default Values (Zero Values) */
	var defaultBool bool
	var defaultInt int
	var defaultFloat float64
	var defaultString string

	fmt.Println("Default bool:", defaultBool)
	fmt.Println("Default int:", defaultInt)
	fmt.Println("Default float64:", defaultFloat)
	fmt.Println("Default string:", defaultString)

	/* structs */
	person := Person{FirstName: "Ritika", LastName: "Agrawal", Age: 23}

	fmt.Printf("person struct is %v", person)
}
