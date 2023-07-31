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

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Legendretransformation")

	/*
		This Go snippet is a simple implementation of the Legendre transformation in the field of thermodynamics. The program represents thermodynamic variables with a struct, Variables, that includes properties like entropy, internal energy, temperature, volume, and pressure. The core function of the program, LegendreTransform, performs the transformation from internal energy representation to the enthalpy representation. Enthalpy is a measure of the total energy of a thermodynamic system and is used to calculate energy changes in a system.The transformation applied in the program is based on the equation H = U + PV, where H is the enthalpy, U is the internal energy, P is the pressure, and V is the volume.

		The Legendre transformation is particularly useful in thermodynamic problems involving processes at nearly constant pressure, as it allows for a change in perspective from internal energy to enthalpy, making the mathematical description of these processes more straightforward. An example illustrating this concept is a gas enclosed in a cylinder with a movable piston (Kolben). As the gas is heated and expands, it does work against the external pressure to push the piston, effectively "making room" for itself. This work, also known as pressure-volume work, is included in the calculation of enthalpy, which is particularly useful for describing systems undergoing changes at nearly constant pressure.

		When the system absorbs heat from its surroundings, its enthalpy increases (ΔH > 0), and the process is said to be endothermic. On the other hand, when the system releases heat to its surroundings, its enthalpy decreases (ΔH < 0), and the process is said to be exothermic.Remember, for isobaric (constant pressure) conditions, the change in enthalpy ΔH directly equals the heat absorbed by the system, i.e., ΔH = q (where q is heat).
	*/

	// define some thermodynamic variables
	v := Variables{
		Entropy:        1.5, // kJ/K
		InternalEnergy: 200, // kJ
		Temperature:    300, // K
		Volume:         1.0, // m^3
		Pressure:       1.0, // bar
	}

	// Perform Legendre transformation
	initialEnthalpy := LegendreTransform(v)

	fmt.Printf("Given Thermodynamic Variables: \n")
	fmt.Printf("Entropy (S): %.2f kJ/K\n", v.Entropy)
	fmt.Printf("Internal Energy (U): %.2f kJ\n", v.InternalEnergy)
	fmt.Printf("Temperature (T): %.2f K\n", v.Temperature)
	fmt.Printf("Volume (V): %.2f m^3\n", v.Volume)
	fmt.Printf("Pressure (P): %.2f bar\n", v.Pressure)
	fmt.Printf("\nApplying Legendre transformation from internal energy (U) to enthalpy (H).\n")
	fmt.Printf("The initial enthalpy of the system (H) is: %.2f kJ\n", initialEnthalpy)

	// Add heat to the system
	const Cp float64 = 1.0    // Specific heat at constant pressure in kJ/(kg*K)
	const deltaT float64 = 50 // Increase in temperature in K
	q := Cp * deltaT          // Heat added in kJ
	v.InternalEnergy += q     // Update internal energy

	// Perform Legendre transformation again after adding heat
	finalEnthalpy := LegendreTransform(v)

	fmt.Printf("\nHeat added to the system (q = Cp*ΔT): %.2f kJ\n", q)
	fmt.Printf("The final enthalpy of the system after adding heat is: %.2f kJ\n", finalEnthalpy)
	fmt.Printf("The change in enthalpy ΔH (which equals heat added at constant pressure) is: %.2f kJ\n", finalEnthalpy-initialEnthalpy)
}
