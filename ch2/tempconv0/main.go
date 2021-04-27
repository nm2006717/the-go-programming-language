package main

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
	symbolF               = "℉"
	symbolC               = "℃"
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string {
	return fmt.Sprintf("%g%s", c, symbolC)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g%s", f, symbolF)
}

func main() {
	//var c = AbsoluteZeroC
	//f := CToF(c)
	//fmt.Printf("%.2f%s = %.2f%s", c, symbolC, f, symbolF)

	fmt.Print("%g\n", BoilingC-FreezingC) //"100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) //"180" °F
	//fmt.Printf("%g\n",boilingF-FreezingC)	// compile error: type mismatch

	var c Celsius
	var f Fahrenheit // "true"
	fmt.Println(c == 0)
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c==f)	// compile error: type mismatch
	fmt.Println(c == Celsius(f)) //"true"

	c = FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}
