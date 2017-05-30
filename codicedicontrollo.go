package codicefiscale

import (
	"regexp"
	"strings"
)

//Codicedicontrollo : rende il codice di controllo del codice fiscale (15 caratteri)
//ingresso: cfin (senza codice di controllo)
//uscita: codice di controllo (vuoto in caso di problemi), err: nil se ok, *CFError se ci sono problemi
func Codicedicontrollo(cfin string) (string, *CFError) {
	if len(cfin) != 15 {
		return "", errCFError("Lunghezza Sbagliata")
	}
	cfin = strings.ToUpper(cfin)
	//verifica per simboli inattesi - usa regexp
	re, _ := regexp.Compile("[^a-zA-Z0-9]")
	if re.MatchString(cfin) {
		return "", errCFError("Caratteri Non Validi")
	}

	s := tcf[string(cfin[14])]
	for i := 0; i <= 13; i += 2 {
		s += tcf[string(cfin[i])] + ordv[string(cfin[i+1])]
	}
	return string(rune(s%26) + rune('A')), nil
}
