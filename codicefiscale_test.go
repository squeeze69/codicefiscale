package codicefiscale

//test file

import (
	"fmt"
	"testing"
)

func TestCodiceFiscale(t *testing.T) {
	//valido
	testOK := []string{"ABCDEF12B23P432P", "MROrSs00a00A000U",""}
	testKO := []string{"ABCDEF12B23P432X", "MRORSS00A00A000V", "MROrSs00a00-A00U", "MRORSS00A.+A000V", "MROrSs00a00A000"}
	for _, v := range testOK {
		s, err := CodiceFiscale(v)
		if !s {
			t.Fatal("Ko: Error! Codice Fiscale", v, " should be valid", err)
		}
		fmt.Printf("Ok (valid) \"%s\"\n", v)
	}
	//codici non validi
	for _, v := range testKO {
		_, err := CodiceFiscale(v)

		if err == nil {
			t.Fatal("Error! Codice Fiscale", v, " should be invalid")
		}
		fmt.Printf("Ok (invalid) \"%s\",%s\n", v,err)
	}
}
