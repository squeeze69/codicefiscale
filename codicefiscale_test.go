package codicefiscale

//test file

import (
	"fmt"
	"testing"
)

func TestCodiceFiscale(t *testing.T) {
	s, err := CodiceFiscale("RSSMCL99M07F205Y")
	fmt.Println("S:", s, "M:", err)
	if !s {
		t.Fatal("Error! Codice Fiscale should be valid")
	}
	s, err = CodiceFiscale("RSSMCL99M07F205Z")
	if s {
		t.Fatal("Error! Codice Fiscale should be invalid!")
	}
	fmt.Println("M:", err)
	s, err = CodiceFiscale("RSSMCL99M07F205")
	if s {
		t.Fatal("Error! Codice Fiscale should be invalid! Short")
	}
	fmt.Println("M:", err)
}
