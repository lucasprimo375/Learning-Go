package auxiliar

import "fmt"

// AuxiliarWrite writes message on the screen informing from which module it was called
func AuxiliarWrite() {
	// public functions in the packaged must have name starting with upper case character
	
	fmt.Println("Writing from auxiliar package")
	secondAuxiliarWrite()
}