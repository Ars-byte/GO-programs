package main

import "fmt"

func main() {
	// Variables are declared with "var" or the ":=" short declaration.
	// The type can be explicit or inferred from the value.
	var explicit int = 10
	inferred := 20

	// Several variables can be declared in one line.
	var a, b string = "hello", "world"
	x, y := 1, 2

	// Constants cannot change and must be known at compile time.
	const answer = 42

	// Basic types:
	//   int      signed integer (platform-sized)
	//   float64  64-bit floating point
	//   string   immutable sequence of UTF-8 bytes
	//   bool     true or false
	//   rune     Unicode code point (alias for int32)
	//   byte     raw byte (alias for uint8)
	var count int = 3
	var ratio float64 = 1.5
	var flag bool = true
	var letter rune = 'A'
	var raw byte = 65

	// Variables without a value get their zero value:
	// 0 for numbers, "" for strings, false for booleans.
	var zeroInt int
	var zeroString string
	var zeroBool bool

	fmt.Println("Variables in Go")
	fmt.Println()
	fmt.Printf("explicit type:      var explicit int = %d\n", explicit)
	fmt.Printf("inferred type:      inferred := %d\n", inferred)
	fmt.Printf("multiple strings:   a, b = %q, %q\n", a, b)
	fmt.Printf("multiple integers:  x, y = %d, %d\n", x, y)
	fmt.Printf("constant:           const answer = %d\n", answer)
	fmt.Printf("int:                %d\n", count)
	fmt.Printf("float64:            %v\n", ratio)
	fmt.Printf("bool:               %v\n", flag)
	fmt.Printf("rune:               %c\n", letter)
	fmt.Printf("byte:               %d\n", raw)
	fmt.Println()
	fmt.Println("Zero values:")
	fmt.Printf("int:    %d\n", zeroInt)
	fmt.Printf("string: %q\n", zeroString)
	fmt.Printf("bool:   %v\n", zeroBool)
	fmt.Println()
	fmt.Println("Type matters: 10/3 =", 10/3, "(int) but 10.0/3 =", 10.0/3, "(float64)")
}
