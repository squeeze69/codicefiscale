package codicefiscale

//test file

import (
	"fmt"
	"testing"
)

func TestCodiceFiscale(t *testing.T) {
	//valido
	s, err := CodiceFiscale("MROrSs00a00A000U")
	fmt.Println("S:", s, "M:", err)
	if !s {
		t.Fatal("Error! Codice Fiscale should be valid")
	}
	//non valido (codice di controllo non corrisponde)
	s, err = CodiceFiscale("MRORSS00A00A000V")
	if s {
		t.Fatal("Error! Codice Fiscale should be invalid!")
	}
	fmt.Println("M:", err)
	//Lunghezza errata
	s, err = CodiceFiscale("MROrSs00a00A000")
	if s {
		t.Fatal("Error! Codice Fiscale should be invalid! Invalid length")
	}
	fmt.Println("M:", err)
}
