/*
1) Gerar o arquivo em Go, com seus métodos.
2) Gerar o .h e o .a através do comando: go build -buildmode=c-archive GoDLL.go
	OBS: Não retirar o comentário //export NomeDaFuncao que está acima das funções, caso contrário não gerará o .h necessário no
	próximo passo.
3) Gerar o .c conforme .c anexo manualmente: ctrl+c e ctrl+v, mudando apenas o nome do .h e o nome do método.
4) Gerar a DLL com o comando: gcc -shared -pthread -o goDLL.dll goDLL.c exportgo.a -lWinMM -lntdll -lWS2_32
5) Criar um arquivo .c conforme anexo e rodar o comando: gcc -o testeGoDLL testeGoDLL.c GoDLL.dll
6) Pronto, agora temos o .exe usando a DLL
*/

package main

import "C"
import "fmt"

//export GetIntFromDLL
func GetIntFromDLL() int32 {
	return 42
}

//export PrintHello
func PrintHello(name string) {
	fmt.Printf("From DLL: Hello, %s!\n", name)
}

//export PrintBye
func PrintBye() {
	fmt.Println("From DLL: Bye!")
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
