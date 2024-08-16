package main

import "fmt"

const ebuliacaoK = 373.0

func main() {

	temperaturaK := ebuliacaoK
	temperaturaC := temperaturaK - 273.0

	fmt.Printf("A temperatura de ebulição da água em Kelvin é %g, e em Celsius é %g", temperaturaK, temperaturaC)
}
