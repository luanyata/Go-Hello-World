package helloworld

import "fmt"

const localeEs = "es"
const localeFr = "fr"
const prefixPtBr = "Olá "
const prefixEs = "Hola "
const prefixFr = "Bonjour "

// Hello Imprime a saudacao de acordo do o idioma selecionado
func Hello(name string, language string) string {
	if name == "" {
		name = "mundo"
	}

	return prefixSalutation(language) + name
}

func prefixSalutation(language string) (prefix string) {
	switch language {
	case localeFr:
		prefix = prefixFr
	case localeEs:
		prefix = prefixEs
	default:
		prefix = prefixPtBr
	}

	return
}

func main() {
	fmt.Println(Hello("mundo", ""))
}
