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

	fmt.Println("test CondiceFiscale")
	//verifica codici validi
	for _, v := range testOK {
		s, err := CodiceFiscale(v)
		if !s {
			t.Fatal("Ko. Errore, Il Codice Fiscale", v, " dovrebbe essere valido", err)
		}
		fmt.Printf("Ok (valido) \"%s\"\n", v)
	}
	//verifica codici non validi
	for _, v := range testKO {
		_, err := CodiceFiscale(v)
		if err == nil {
			t.Fatal("Ko. Errore, Il Codice Fiscale", v, " NON dovrebbe essere valido")
		}
		fmt.Printf("Ok (non valido) \"%s\",%s\n", v, err)
	}
}

func TestCodicedicontrollo(t *testing.T) {
	fmt.Println("test Codicedicontrollo")
	if s, err := Codicedicontrollo("ABCDEF12B23P432"); s != "P" {
		t.Fatal("Ko. Errore, il codice di controllo è sbagliato atteso \"P\", avuto", s, err)
	}
	fmt.Println("Ok (valido) codice di controllo corrisponde")
}

func TestConfrontaCodicifiscaliOmocodici(t *testing.T) {
	fmt.Println("test ConfrontaCodicifiscaliOmocodici")
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

func TestConfrontaCodicifiscali(t *testing.T) {
	fmt.Println("test ConfrontaCodicifiscali")
	if _, err := ConfrontaCodicifiscali("ABCDEF12B23P432P", "ABCDEF12B23P432P"); err != nil {
		t.Fatal("KO. Errore, dovrebbe essere uguale", err)
	}
	fmt.Println("Ok. Uguali")

	a := "ABCDEF12B23P432P"
	b := "ABCDEF12B23P433R"
	if _, err := ConfrontaCodicifiscali(a, b); err == nil {
		t.Fatal("KO. Errore, dovrebbero essere diversi", a, b, err)
	}
	fmt.Println("Ok. Diversi")
}

// test aggiuntivi utili per godoc
func ExampleConfrontaCodicifiscali() {
	if _, err := ConfrontaCodicifiscali("ABCDEF12B23P432P", "ABCDEF12B23P433R"); err != nil {
		fmt.Println("ConfrontaCodicifiscali:", err)
	}
	// Output: ConfrontaCodicifiscali: Non corrispondono
}

func ExampleCodicedicontrollo() {
	if p, err := Codicedicontrollo("ABCDEF12B23P433"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p)
	}
	// Output: R
}

func ExampleCodiceFiscale() {
	if _, err := CodiceFiscale("ABCDEF12B23P433R"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Codice Fiscale Verificato")
	}
	// Output: Codice Fiscale Verificato
}
