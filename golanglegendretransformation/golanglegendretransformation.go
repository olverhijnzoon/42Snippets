package main

import (
	"fmt"
)

// Variables represent a set of thermodynamic variables
type Variables struct {
	Entropy        float64 // S
	InternalEnergy float64 // U
	Temperature    float64 // T
	Volume         float64 // V
	Pressure       float64 // P
}

// LegendreTransform performs a Legendre Transformation on the given variables
// This particular example implements the transformation from the internal energy representation (U) to the enthalpy representation (H)
func LegendreTransform(v Variables) float64 {
	return v.InternalEnergy + v.Pressure*v.Volume
}

func main() {
	// define some thermodynamic variables
	v := Variables{
		Entropy:        1.5, // kJ/K
		InternalEnergy: 200, // kJ
		Temperature:    300, // K
		Volume:         1.0, // m^3
		Pressure:       1.0, // bar
	}

	// perform Legendre transformation
	result := LegendreTransform(v)

	fmt.Printf("Given Thermodynamic Variables: \n")
	fmt.Printf("Entropy (S): %.2f kJ/K\n", v.Entropy)
	fmt.Printf("Internal Energy (U): %.2f kJ\n", v.InternalEnergy)
	fmt.Printf("Temperature (T): %.2f K\n", v.Temperature)
	fmt.Printf("Volume (V): %.2f m^3\n", v.Volume)
	fmt.Printf("Pressure (P): %.2f bar\n", v.Pressure)
	fmt.Printf("\nApplying Legendre transformation from internal energy (U) to enthalpy (H).\n")
	fmt.Printf("The result of the transformation (H = U + PV) is: %.2f kJ\n", result)
}
