package main

import (
	"monitoring/internal"
)

func main() {
/*
	// almost every return value is a struct
	c, _ := cpu.Info()
	t, _ := cpu.Percent(10*time.Millisecond, false)
	indent, _ := json.MarshalIndent(c[0], "", "\t")
	istring := string(indent)
	times, _ := cpu.Times(true)

	// convert to JSON. String() is also implemented
	fmt.Println(istring)
	fmt.Println(t)
	fmt.Println(times)*/

	internal.SetConfigFile()
}