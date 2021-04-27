package main

import (
	"./lengthconv"
	"./tempconv"
	"./weightconv"
	"flag"
	"fmt"
)

var w = flag.Float64("w", 0, "重量转换")
var t = flag.Float64("t", 0, "温度转换")
var l = flag.Float64("l", 0, "长度转换")

func main() {

	flag.Parse()

	if !isFlagPassed("w") {
		fmt.Println("hhhhhh")
	}

	if *l != 0 {
		fmt.Printf("%g%s = %g%s\t\t", *l, lengthconv.SymbolM, lengthconv.MtoF(lengthconv.Meter(*l)), lengthconv.SymbolFt)
		fmt.Printf("%g%s = %g%s\n", *l, lengthconv.SymbolFt, lengthconv.FToM(lengthconv.Foot(*l)), lengthconv.SymbolM)
	}

	if *w != 0 {
		fmt.Printf("%g%s = %g%s\t\t", *w, weightconv.SymbolKg, weightconv.KgToLb(weightconv.Kilogram(*w)), weightconv.SymbolLb)
		fmt.Printf("%g%s = %g%s\n", *w, weightconv.SymbolLb, weightconv.LbToKg(weightconv.Pound(*w)), weightconv.SymbolKg)
	}
	if *t != 0 {
		fmt.Printf("%g%s = %g%s\t\t", *t, tempconv.SymbolF, tempconv.FToC(tempconv.Fahrenheit(*t)), tempconv.SymbolC)
		fmt.Printf("%g%s = %g%s\n", *t, tempconv.SymbolC, tempconv.CToF(tempconv.Celsius(*t)), tempconv.SymbolC)

		fmt.Printf("%g%s = %g%s\t\t", *t, tempconv.SymbolF, tempconv.FToK(tempconv.Fahrenheit(*t)), tempconv.SymbolK)
		fmt.Printf("%g%s = %g%s\n", *t, tempconv.SymbolK, tempconv.KToF(tempconv.Kelvin(*t)), tempconv.SymbolF)

		fmt.Printf("%g%s = %g%s\t\t", *t, tempconv.SymbolC, tempconv.CToK(tempconv.Celsius(*t)), tempconv.SymbolK)
		fmt.Printf("%g%s = %g%s\n", *t, tempconv.SymbolK, tempconv.KToC(tempconv.Kelvin(*t)), tempconv.SymbolC)

	}

}

// 判断用户是否输入参数。
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
