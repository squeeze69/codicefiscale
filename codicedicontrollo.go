package codicefiscale

import (
	"regexp"
	"strings"
)

/*
Genera codice di controllo codice fiscale
Versione: 1.0
Data: 1/5/2017
Autore: Squeeze69
Licenza: LGPL
Porting basato sulle informazioni pubblicate da Umberto Salsi su Icosaedro:
sito web: http://www.icosaedro.it/cf-pi/index.html

package: https://github.com/squeeze69/codicefiscale
con go: go get github.com/squeeze69/codicefiscale
*/

//Codicedicontrollo : rende il codice di controllo del codice fiscale (15 caratteri)
//ingresso: cfin (senza codice di controllo)
//uscita: codice di controllo (vuoto in caso di problemi), err: nil se ok, *CFError se ci sono problemi
func Codicedicontrollo(cfin string) (string, *CFError) {
	if len(cfin) != 15 {
		return "", errCFError("Lunghezza Sbagliata")
	}
	//verifica per simboli inattesi
	if regexp.MustCompile("[^a-zA-Z0-9]").MatchString(cfin) {
		return "", errCFError("Caratteri Non Validi")
	}
	cfin = strings.ToUpper(cfin)
	s := tcf[cfin[14]]
	for i := 0; i <= 13; i += 2 {
		s += tcf[cfin[i]] + ordv[cfin[i+1]]
	}
	return string(rune(s%26) + rune('A')), nil
}
