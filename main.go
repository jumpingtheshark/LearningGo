package main // you have to have this package main here otehrwise the thing doesn't run!!!!
import "math/rand"

// dumb

type kelvin float64


func fakeSensor() kelvin {
	return kelvin (rand.Intn (151 + 150))

}


func realSensor() kelvin {
	return 0

}


func main (){

	sensor:= fakeSensor

	sensor()


}
