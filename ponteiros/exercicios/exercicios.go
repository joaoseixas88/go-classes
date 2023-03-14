package main

import "fmt"

// Crie um valor e atribua-o a uma variavel
// Demonstre o endereco na memoria

func main() {
	exercicio_1()
	person := Person{
		name: "Joao",
		age:  25,
	}
	person.mudeMe()
	exercicio_2(person)
	fmt.Println(person)
}

func exercicio_1() {
	x := 10
	fmt.Println(&x)

}

// - Crie um struct "pessoa"
// - Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
//     - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
//     - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
//     - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
//     - → x.f is shorthand for (*x).f." ←
//     - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
// - Crie um valor do tipo pessoa;
// - Use a função mudeMe e demonstre o resultado.

type Person struct {
	name string
	age  int
}

func (p *Person) mudeMe() {
	(*p).age = 33
}

func exercicio_2(p Person) {
	p.mudeMe()
}
