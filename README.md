# Verifica Codice Fiscale in [GO](http://golang.org)

## Licenza: LGPLv3

**Package**: github.com/squeeze69/codicefiscale

**Download del package**: go get github.com/squeeze69/codicefiscale

Porting basato sulle informazioni pubblicate da Umberto Salsi su [Icosaedro](http://www.icosaedro.it/cf-pi/index.html)

```
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
```

Simile ma per la [partita IVA](https://github.com/squeeze69/partitaiva)
