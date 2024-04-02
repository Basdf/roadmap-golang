package basics //nombre del paquete al cual pertenece la clase

import (
	"fmt"

	"rsc.io/quote" //modulos importados
)

func FmtSayHello() {
	fmt.Println("Hello World!")
}

func RSCQuoteSayGo() {
	fmt.Println(quote.Go())
}
