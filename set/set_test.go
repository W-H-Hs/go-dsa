package set

import (
	"fmt"
	"testing"
)

type IAnimal interface {
	Run()
}

type Animal struct {
	Name string
	Age  int
	IAnimal
}

type People struct {
	Animal
}

func (p People) Run() {
	fmt.Println("run")
}

func testFun(animal IAnimal) {
	fmt.Println(animal.(Animal).Name)
	fmt.Println(animal.(Animal).Age)
}

func Test(t *testing.T) {
	p := People{Animal{
		Name: "xiaoming",
		Age:  18,
	}}
	p.IAnimal = p
	testFun(p)
}
