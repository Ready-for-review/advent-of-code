package main

import "fmt"

type ModuleMass int

func (m ModuleMass) String() string {
	return fmt.Sprintf("%d", m)
}

type Fuel int

func (f Fuel) String() string {
	return fmt.Sprintf("%d", f)
}

// for Stefan: func CalculateFuel(moduleMass int) int
func CalculateFuel(moduleMass ModuleMass) Fuel {
	return Fuel(moduleMass/3 - 2)
}
