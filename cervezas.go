package main

import "fmt"
import "errors"

/*
	-Elegir un dominio (ej, autos, cervezas, mascotas…) que prefieran
	-Implementar una “agencia de autos” que guarda internamente “autos” (adaptado a su contexto)
	-La “agencia” debe permitir crear, leer, actualizar y eliminar autos
*/

type Printable interface {
	print()
}

type cerveza struct {
	name string
}

func (b *cerveza) print(){
	fmt.Println(b.name)
}

func (b *cerveza) equals(n string)(int, error){
	if b.name == n{
		return 1, nil
	}
	return 0, errors.New("Son diferentes")
}

type cerveceria struct {
	name string
	l []cerveza
	//var l []string
	//l:= make([]string, 10)
	//l []string
}

func (c *cerveceria) crearCerveza(n string){
	beer := &cerveza{name: n}
	c.l =append(c.l, beer)
}

func (c *cerveceria) print(){
	for i := range c.l {
		c.l[i].print()
	}
}

func (c*cerveceria) eliminarCerveza(n string){
	for i := range c.l {
		r, err := c.l[i].equals(n) 
		fmt.Println(c.l[i])
		if err == nil {
			//c.l = delete(c.l[i], v)
		}
	}
}

func main(){
	c := &cerveceria{name: "Brewery"}
	c.crearCerveza("IPA")
	c.crearCerveza("Porter")
	c.crearCerveza("Scotch")
	c.crearCerveza("Scotish")
	c.crearCerveza("Honney")
	c.crearCerveza("APA")
	c.print()
	c.eliminarCerveza("Honney")
	c.print()
}