package main

type kelvin float64


type sensor func() kelvin

func realSensor() kelvin{

	return 0;
}



func calibrate (s sensor, offset kelvin) sensor {

	return func() kelvin {
		return s() + offset
	}

}

func main (){

	// so we effectively pick which function to pass in
	sensor := calibrate( realSensor, 5)
	print(sensor())


}