package codicefiscale

//test file

import (
	"fmt"
	"testing"
)

func TestCodiceFiscale(t *testing.T) {
	//codici validi (formalmente - ogni somiglianza con eventuali codici reali è casuale)
	testOK := []string{"ABCDEF12B23P432P", "MROrSs00a00A000U", ""}
	//codici non validi
	testKO := []string{"ABCDEF12B23P432X", "MRORSS00A00A000V", "MROrSs00a00-A00U", "MRORSS00A.+A000V", "MROrSs00a00A000"}
	for _, v := range testOK {
		s, err := CodiceFiscale(v)
		if !s {
			t.Fatal("Ko. Errore, Il Codice Fiscale", v, " dovrebbe essere valido", err)
		}
		fmt.Printf("Ok (valido) \"%s\"\n", v)
	}
	//codici non validi
	for _, v := range testKO {
		_, err := CodiceFiscale(v)

		if err == nil {
			t.Fatal("Ko. Errore, Il Codice Fiscale", v, " NON dovrebbe essere valido")
		}
		fmt.Printf("Ok (non valido) \"%s\",%s\n", v, err)
	}
}
