package main

import "fmt"
import "errors"

//Si la funcion empieza con minúscula se puede usar desde otro package pero no
//desde otro directorio (es privada). Si comienza con mayúscul es pública y hay 
//que documentarlo.
func suma(a,b int)(int,error){
	if a < b {
		return 0, errors.New("el primer valor es menor que el segundo valor")
	}
	return a+b,nil
}
/* otra manerade hacer la misma funcion
func suma(a,b int)(c int,e error){
	c= a+b
	e=nil
	if a < b {
		e = errors.New("el primer valor es menor que el segundo valor")
		c=0
	}
	return
}
*/
	

func main(){
	/*
	fmt.Println("Hello world")
	r, err := suma(1,2)
	if err != nil {
		fmt.Println(err)
		} else {
			fmt.Println(r)
		}
	}
	fmt.Println("------------------")
*/
	fmt.Println("ARRAY")
	//DECLARAR ARRAY
	var arr [2]int 
	arr[0] = 1
	arr[1] = 2
	
	//RECORRER EL ARRAY
	for i,v := range arr{
		fmt.Println(i, v)
	}
	fmt.Println("------------------")
	fmt.Println("SLICES")
	var l []int
	//l:= make([]int, 100) //a l lo vuelve un slice con capacidad 100, silos supero, me agrega 100+
	l= append(l, 10)
	fmt.Printf("%p\n", l) //%p es la posicion de memoria
	l= append(l, 100)
	fmt.Printf("%p\n", l) //Printf es el formato en el que se vaa mostrar
	l= append(l, 1000)
	fmt.Printf("%p\n", l)
	for i,v := range l {
		fmt.Println(i, v)
	}

	fmt.Println("------------------")
	fmt.Println("MAPA")
	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"
	for k,v := range m{
		fmt.Println(k, v)
	}
	
}