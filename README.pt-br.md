[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-closure-function/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-closure-function/blob/main/README.md)  
go version 1.22.1  

## Função de fechamento (Closure function)
São funções que referênciam variáveis que estão fora do seu corpo/escopo:  
```go
package main

import "fmt"

func closure() func() { //retorna uma função
	text := "Dentro da função closure"
	function := func() { // só printa um texto na tela
		fmt.Println(text)
	}
	return function
}

func main() {
	text := "Dentro da função main"
	fmt.Println(text)

	funcN := closure()
	funcN()
}
```
Agora vamos pensar no que vai acontecer aqui, em ```closure()``` estamos retornando uma função sem nome e sem parâmetro que printa um texto, mas quando ela for chamada na ```main()```, ele irá imprimir o texto da função ou o texto que está na main:  
```go
// outputs
  fmt.Println(text) // Dentro da função main
  funcN := closure()
  funcN() // Dentro da função closure
```
Isso é um comportamento diferente da função ```closure()```, porquê quando criamos a função estávemos referênciando a variável do escopo ```text```, portanto quando executarmos ela, essa referência não sera perdida, ela mantem a referência inicial.  
A função ```closure``` é isso, ela tem uma espécie de "memória", de onde veio, e sempre estaremos no referindo as referências passadas à ela na função.  