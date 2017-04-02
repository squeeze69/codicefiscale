package codicefiscale

//CFError - error structure for CodiceFiscale
type CFError struct {
	msg string
}

func (cf *CFError) Error() string {
	return cf.msg
}

//CodiceFiscale  controlla il codice fiscale, restituisce doppio valore, true/false e messaggio ove opportuno
func CodiceFiscale(cfin string) (bool, *CFError) {
	if len(cfin) != 16 {
		er := new(CFError)
		er.msg = "Lunghezza Errata"
		return false, er
	}
	//decodifica carattere di controllo cf
	tcf := map[string]int{
		"0": 1, "1": 0, "2": 5, "3": 7, "4": 9, "5": 13, "6": 15, "7": 17, "8": 19,
		"9": 21, "A": 1, "B": 0, "C": 5, "D": 7, "E": 9, "F": 13, "G": 15, "H": 17,
		"I": 19, "J": 21, "K": 2, "L": 4, "M": 18, "N": 20, "O": 11, "P": 3, "Q": 6, "R": 8,
		"S": 12, "T": 14, "U": 16, "V": 10, "W": 22, "X": 25, "Y": 24, "Z": 23,
	}
	//ordinamento
	ordv := map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7,
		"I": 8, "J": 9, "K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17,
		"S": 18, "T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25,
	}
	var s int
	for i := 1; i <= 13; i += 2 {
		s += ordv[string(cfin[i])]
	}
	for i := 0; i <= 14; i += 2 {
		s += tcf[string(cfin[i])]
	}
	if s%26 != ordv[string(cfin[15])] {
		er := new(CFError)
		er.msg = "Codice Di Controllo"
		return false, er
	}
	return true, nil
}
