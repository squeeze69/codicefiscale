package codicefiscale

import "strings"

//	caratteri da modificare, base 1: 7, 8, 10, 11, 13, 14, 15
var omcndx = map[int]bool{
	6: true, 7: true, 9: true, 10: true, 12: true, 13: true, 14: true,
}

var omc = map[string]string{
	"L": "0", "M": "1", "N": "2", "P": "3", "Q": "4",
	"R": "5", "S": "6", "T": "7", "U": "8", "V": "9",
}

//inverte le variazioni per omocodie
func deomocodia(s string) string {
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

//ConfrontaCodicifiscali : ingresso a,b, tiene conto di omocodie per confronto
//ingresso: a,b stringhe con i codici fiscali da confrontare
//se non corrispondono, riconduce entrambi alla forma non per omocodie e riconfronta
//uscita: bool (true:ok,false:ko), *CFError (nil se va bene)
//da sostituire: 7,8,10,11,13,14,15
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
	ad := deomocodia(a)
	if strings.Compare(ad[0:15], b[0:15]) != 0 {
		er := new(CFError)
		er.msg = "Non corrispondono"
		return false, er
	}
	return true, nil
}
