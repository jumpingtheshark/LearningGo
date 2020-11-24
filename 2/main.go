package main

import "fmt"

type kelvin float64

// the only thing that I am getting out of this story is that the anonymous functino that sensor pionts to
// has access to
func main (){
	var k kelvin = 294.0
	sensor := func() kelvin {
		return k
	}

 	fmt.Println(sensor())
	k++

	fmt.Println(sensor())


}