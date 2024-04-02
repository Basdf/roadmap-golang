package basics //nombre del paquete al cual pertenece la clase

import (
	"fmt"

	"rsc.io/quote" //modulos importados
)

func fmtSayHello() {
	fmt.Println("Hello World!")
}

func rscQuoteSayGo() {
	fmt.Println(quote.Go())
}
