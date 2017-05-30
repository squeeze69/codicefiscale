package codicefiscale

import (
	"regexp"
	"strings"
)

/*
Verifica codice fiscale
Versione: 1.0
Data: 1/5/2017
Autore: Squeeze69
Licenza: LGPL
Porting basato sulle informazioni pubblicate da Umberto Salsi su Icosaedro:
sito web: http://www.icosaedro.it/cf-pi/index.html

package: https://github.com/squeeze69/codicefiscale
con go: go get github.com/squeeze69/codicefiscale
*/

//CFError - error structure for CodiceFiscale - identificabile con typecast
type CFError struct {
	msg string
}

func (cf *CFError) Error() string {
	return cf.msg
}

//decodifica carattere di controllo cf
var tcf = map[string]int{
	"0": 1, "1": 0, "2": 5, "3": 7, "4": 9, "5": 13, "6": 15, "7": 17, "8": 19,
	"9": 21, "A": 1, "B": 0, "C": 5, "D": 7, "E": 9, "F": 13, "G": 15, "H": 17,
	"I": 19, "J": 21, "K": 2, "L": 4, "M": 18, "N": 20, "O": 11, "P": 3, "Q": 6, "R": 8,
	"S": 12, "T": 14, "U": 16, "V": 10, "W": 22, "X": 25, "Y": 24, "Z": 23,
}

//map per simulare "ord" di altri linguaggi - più semplice di int(rune) - int('A') oppure int('0')
var ordv = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7, "I": 8, "J": 9,
	"K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17, "S": 18,
	"T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25,
}

//genera un errore CFError
func errCFError(s string) *CFError {
	er := new(CFError)
	er.msg = s
	return er
}

//CodiceFiscale  controlla il codice fiscale, restituisce doppio valore, true/false e messaggio, ove opportuno
//se cfin è vuota, viene considerata valida, per questo caso, il controllo dovrebbe essere altrove
//Ingresso: cfin: stringa,non importa maiuscolo o minuscolo
//Uscita: bool:true (a posto)/false (problemi) e *CFError (nil (a posto)/puntatore all'errore (problemi)
func CodiceFiscale(cfin string) (bool, *CFError) {

	if len(cfin) == 0 {
		return true, nil //convenzione generale usata sulle routine su Icosaedro
	}
	if len(cfin) != 16 {
		return false, errCFError("Lunghezza Sbagliata")
	}

	//verifica per simboli inattesi - usa regexp
	re, _ := regexp.Compile("[^a-zA-Z0-9]")
	if re.MatchString(cfin) {
		return false, errCFError("Caratteri Non validi")
	}

	cfin = strings.ToUpper(cfin)
	s := tcf[string(cfin[14])]
	for i := 0; i <= 13; i += 2 {
		s += tcf[string(cfin[i])] + ordv[string(cfin[i+1])]
	}
	if s%26 != ordv[string(cfin[15])] {
		return false, errCFError("Carattere Di Controllo Non Valido")
	}
	return true, nil
}
