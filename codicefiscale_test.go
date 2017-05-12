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

func TestCodicedicontrollo(t *testing.T) {
	if s, err := Codicedicontrollo("ABCDEF12B23P432"); s != "P" {
		t.Fatal("Ko. Errore, il codice di controllo è sbagliato atteso \"P\", avuto", s, err)
	}
	fmt.Println("Ok (valido) codice di controllo corrisponde")
}

func TestConfrontaCodicifiscaliOmocodici(t *testing.T) {
	s, _ := Codicedicontrollo("ABCDEF12B23P432")
	o, _ := Codicedicontrollo("ABCDEF12B23P43N")
	sb := "ABCDEF12B23P432" + s
	oa := "ABCDEF12B23P43N" + o
	if _, err := ConfrontaCodicifiscaliOmocodici(oa, sb); err != nil {
		t.Fatal("KO. Errore, dovrebbe essere uguale", oa, sb, err)
	}
	fmt.Println("Ok. Uguale - test anche con omocodie ", oa, sb)

	s, _ = Codicedicontrollo("ABCDEF12B23P433")
	o, _ = Codicedicontrollo("ABCDEF12B23P43N")
	sb = "ABCDEF12B23P433" + s
	oa = "ABCDEF12B23P43N" + o
	if _, err := ConfrontaCodicifiscaliOmocodici(oa, sb); err == nil {
		t.Fatal("KO. Errore, dovrebbero essere diversi", oa, sb, err)
	}
}
