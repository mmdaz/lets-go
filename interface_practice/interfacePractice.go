package interface_practice

import "fmt"

type InterfacePractice interface {
	RunInterfacePractice(index int)
}

type InterfacePracticeImp struct {

}

func NewInterfacePractice() InterfacePractice {
	return &InterfacePracticeImp{}
}

func (i *InterfacePracticeImp) RunInterfacePractice(index int) {
	fmt.Println("interface practice: ", index)
}
