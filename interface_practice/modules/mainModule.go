package modules

import "fmt"

type MainModule interface {
	RunMainModule(index int)
}

type MainModuleImp struct {

}

func NewMainModule() MainModule {
	return &MainModuleImp{}
}

func (i *MainModuleImp) RunMainModule(index int) {
	fmt.Println("interface practice: ", index)
}

func (d *MainModuleImp) innerFunction() {
	panic("implement me!")
}