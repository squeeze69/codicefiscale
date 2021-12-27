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
		err := CodiceFiscale(v)
		if err != nil {
			t.Fatal("Ko. Errore, Il Codice Fiscale", v, " dovrebbe essere valido", err)
		}
		fmt.Printf("Ok (valido) \"%s\"\n", v)
	}
	//verifica codici non validi
	for _, v := range testKO {
		err := CodiceFiscale(v)
		if err == nil {
			t.Fatal("Ko. Errore, Il Codice Fiscale", v, " NON dovrebbe essere valido")
		}
		fmt.Printf("Ok (non valido) \"%s\",%s\n", v, err)
	}
}

func TestCodicedicontrollo(t *testing.T) {
	var s string
	var err error
	fmt.Println("test Codicedicontrollo")
	if s, err = Codicedicontrollo("ABCDEF12B23P432"); s != "P" {
		t.Fatal("Ko. Errore, il codice di controllo è sbagliato atteso \"P\", avuto", s, err)
	}
	fmt.Println("Ok (valido) codice di controllo corrisponde")
	if _, err = Codicedicontrollo("ABCDEF12B23P43"); err.(*CFError) == nil {
		t.Fatal("Ko. Errore, la lunghezza è sbagliata, dovrebbe restituire errore")
	}
	fmt.Println("Ok (non valido) lunghezza sbagliata")
	if _, err = Codicedicontrollo("ABCDEF12B23P*32"); err.(*CFError) == nil {
		t.Fatal("Ko. Errore, carattere non ammesso, dovrebbe restituire errore")
	}
	fmt.Println("Ok (non valido) carattere non ammesso")
}

func TestConfrontaCodicifiscaliOmocodici(t *testing.T) {
	const (
		cof1 = "ABCDEF12B23P432"
		cof2 = "ABCDEF12B23P43N"
		cof3 = "ABCDEF12B23P432P"
	)
	fmt.Println("test ConfrontaCodicifiscaliOmocodici")
	s, _ := Codicedicontrollo(cof1)
	o, _ := Codicedicontrollo(cof2)
	sb := cof1 + s
	oa := cof2 + o
	if err := ConfrontaCodicifiscaliOmocodici(oa, sb); err != nil {
		t.Fatal("KO. Errore, dovrebbe essere uguale", oa, sb, err)
	}
	fmt.Println("Ok. Uguale - test anche con omocodie ", oa, sb)
	s, _ = Codicedicontrollo("ABCDEF12B23P433")
	sb = "ABCDEF12B23P433" + s
	if err := ConfrontaCodicifiscaliOmocodici(oa, sb); err == nil {
		t.Fatal("KO. Errore, dovrebbero essere diversi", oa, sb)
	}
	if err := ConfrontaCodicifiscaliOmocodici(cof1, sb); err == nil {
		t.Fatal("KO. Errore, dovrebbe essere uguale", oa, sb)
	}
	if err := ConfrontaCodicifiscaliOmocodici(oa, cof2); err == nil {
		t.Fatal("KO. Errore, dovrebbe essere uguale", oa, sb)
	}
	if err := ConfrontaCodicifiscaliOmocodici(cof3, cof3); err != nil {
		t.Fatal("KO. Errore, dovrebbe essere uguale", cof3, err)
	}
	fmt.Println("Ok, errori per codici fiscali sbagliati e giusti")
}

func TestConfrontaCodicifiscali(t *testing.T) {
	const (
		cof1sbagliato = "ABCDEF12B23P432X"
		cofa          = "ABCDEF12B23P432P"
		cofb          = "ABCDEF12B23P433R"
	)
	fmt.Println("test ConfrontaCodicifiscali")
	if err := ConfrontaCodicifiscali(cofa, cofa); err != nil {
		t.Fatal("KO. Errore, dovrebbe essere uguale", err)
	}
	fmt.Println("Ok. Uguali")

	if err := ConfrontaCodicifiscali(cofa, cofb); err == nil {
		t.Fatal("KO. Errore, dovrebbero essere diversi", cofa, cofb, err)
	}
	fmt.Println("Ok. Diversi")
	if err := ConfrontaCodicifiscali(cof1sbagliato, cofb); err == nil {
		t.Fatal("KO. Errore, dovrebbe rilevare codice sbagliato", cof1sbagliato, err)
	}
	if err := ConfrontaCodicifiscali(cofa, cof1sbagliato); err == nil {
		t.Fatal("KO. Errore, dovrebbe rilevare codice sbagliato", cof1sbagliato, err)
	}
	fmt.Println("Ok. Rileva codici sbagliati")
}

// test aggiuntivi utili per godoc
func ExampleConfrontaCodicifiscali() {
	if err := ConfrontaCodicifiscali("ABCDEF12B23P432P", "ABCDEF12B23P433R"); err != nil {
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
	if err := CodiceFiscale("ABCDEF12B23P433R"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Codice Fiscale Verificato")
	}
	// Output: Codice Fiscale Verificato
}
