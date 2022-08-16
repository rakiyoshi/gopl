// tempconv package calculates temperatures of the Celsius and Fahrenheit
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	FreezingC Celsius = 0
	BoilingC  Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)

}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	// compile error
	// fmt.Printf("%g\n", boilingF-FreezingC)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// compile error
	// fmt.Println(c == f)
	fmt.Println(c == Celsius(f))

	c = FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"
	fmt.Println(float64(c)) // "100"
}
