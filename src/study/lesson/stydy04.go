package lesson

import "fmt"

type person struct {
	name string
	age int
	sex string
	skills
}


type skills []string

func Four(){
	P := person{name:"Tom",age:25}
	P.skills = []string{"1","3","5"}
	P.sex = "gril"
	fmt.Println(P,P.name)
}
