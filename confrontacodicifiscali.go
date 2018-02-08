package codicefiscale

import "strings"

/*
Confronto Codici Fiscali
Versione: 1.0
Data: 1/5/2017
Autore: Squeeze69
Licenza: LGPL
Porting basato sulle informazioni pubblicate da Umberto Salsi su Icosaedro:
sito web: http://www.icosaedro.it/cf-pi/index.html

package: https://github.com/squeeze69/codicefiscale
con go: go get github.com/squeeze69/codicefiscale
*/

//	caratteri da modificare, base 1: 7, 8, 10, 11, 13, 14, 15
var omcndx = map[int]bool{
	6: true, 7: true, 9: true, 10: true, 12: true, 13: true, 14: true,
}

// mappatura per invertire trasformazioni in caso di omocodie
var omc = map[string]string{
	"L": "0", "M": "1", "N": "2", "P": "3", "Q": "4",
	"R": "5", "S": "6", "T": "7", "U": "8", "V": "9",
}

//Deomocodia inverte le variazioni per omocodie, assume un codice fiscale valido, tutte le lettere maiuscole
// vocali accentate convertite con vocali non accentate
func Deomocodia(s string) string {
	var s2 string
	for i, c := range s[0:15] {
		if _, ok := omcndx[i]; ok {
			if v, vb := omc[string(c)]; vb {
				s2 = s2 + v
			} else {
				s2 = s2 + string(c)
			}
		} else {
			s2 = s2 + string(c)
		}
	}
	return s2
}

//ConfrontaCodicifiscaliOmocodici : ingresso a,b, tiene conto di omocodie per confronto
//ingresso: a,b stringhe con i codici fiscali da confrontare
//se non corrispondono, riconduce entrambi alla forma non per omocodie e riconfronta
//uscita: bool (true:ok,false:ko), *CFError (nil se va bene)
//da sostituire: 7,8,10,11,13,14,15
func ConfrontaCodicifiscaliOmocodici(a, b string) (bool, *CFError) {
	if _, err := CodiceFiscale(a); err != nil {
		return false, err
	}
	if _, err := CodiceFiscale(b); err != nil {
		return false, err
	}
	a = strings.ToUpper(a)
	b = strings.ToUpper(b)
	if strings.Compare(a, b) == 0 {
		return true, nil
	}
	ad := Deomocodia(a)
	bd := Deomocodia(b)
	if strings.Compare(ad[0:15], bd[0:15]) != 0 {
		return false, errCFError("Non corrispondono")
	}
	return true, nil
}

//ConfrontaCodicifiscali : controlla e confronta due codici fiscali
//DEVONO corrispondere al 100% - prima verifica il codice di controllo
//Ingresso: a,b : string : codifici fiscali
//Uscita: bool (true:ok,false:ko), *CFError (nil se va bene)
func ConfrontaCodicifiscali(a, b string) (bool, *CFError) {
	if _, err := CodiceFiscale(a); err != nil {
		return false, err
	}
	if _, err := CodiceFiscale(b); err != nil {
		return false, err
	}
	a = strings.ToUpper(a)
	b = strings.ToUpper(b)
	if strings.Compare(a, b) == 0 {
		return true, nil
	}
	return false, errCFError("Non corrispondono")
}
