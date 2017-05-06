Verifica Codice Fiscale in GO (http://golang.org)

License: LGPLv3

2017

package: go get github.com/squeeze69/codicefiscale

Porting basato sulle informazioni pubblicate da Umberto Salsi su Icosaedro:
sito web: http://www.icosaedro.it/cf-pi/index.html

package main

import (
	"github.com/squeeze69/codicefiscale"
	"fmt"
)

func main() {
	ok, err := codicefiscale.CodiceFiscale("ABCDEF12B23P432P")
	if err != nil {
		fmt.Println("Codice Fiscale non valido:",err)
	} else {
		fmt.Println("Codice Fiscale Valido")
	}
}

per la partita IVA, vedi: https://github.com/squeeze69/partitaiva
