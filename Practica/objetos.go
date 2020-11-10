package main

import "fmt"

type Printable interface {
	print()
}
/*
	La única manera de que un objeto implemente una iterface es que elobjeto 
	implemente TODOS los métodos de la interface
*/

type person struct {
	name string
}

type figure struct {
	name string
}

func (f *figure) print(){
	fmt.Println("{figure}", f.name)
}

func invokePrint(p Printable){
	p.print()
}

func (p *person) print(){
	fmt.Println("{person}", p.name)
}

func (p *person) clean(){
	p.name = ""
}

func main(){
	p := &person{name: "Juan"}
	invokePrint(p)

	f := &figure{name: "Cube"}
	invokePrint(f)
}