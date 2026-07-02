package main

import "fmt"

func main() {
	// Signed Integer
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	// Unsigned Integer
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	fmt.Println("===== Signed Integer =====")
	fmt.Println("int8  :", i8)
	fmt.Println("int16 :", i16)
	fmt.Println("int32 :", i32)
	fmt.Println("int64 :", i64)

	fmt.Println()

	fmt.Println("===== Unsigned Integer =====")
	fmt.Println("uint8  :", u8)
	fmt.Println("uint16 :", u16)
	fmt.Println("uint32 :", u32)
	fmt.Println("uint64 :", u64)
}