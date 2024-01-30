package interfaces

import "fmt"

const legalAge uint8 = 18

type Human interface {
	breath() bool
}

type Man struct {
	name string
	age  uint8
}

type Woman struct {
	name string
	age  uint8
}

func (man Man) canDrink() bool {
	return man.age > legalAge
}
func (man Man) breath() bool {
	return false
}
func (woman Woman) breath() bool {
	return true
}

func UseInterfaces() {
	man := Man{name: "rakesh", age: 20}
	woman := Woman{name: "jasmeet", age: 30}
	fmt.Println(man.canDrink())
	fmt.Println(isAlive(man))
	fmt.Println(man)

	fmt.Println(woman.breath())
	fmt.Println(isAlive(woman))
	fmt.Println(woman)
}

func isAlive(human Human) string {
	checking := fmt.Sprintf("%t this is something", human.breath())
	return checking
}

// we have decoupled the isAlive method from man or woman... both of them can use the method.
