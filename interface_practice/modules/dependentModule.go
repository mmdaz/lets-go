package modules

import "fmt"

type DependentModule interface {
	DoPractice(printableString string)

}

type DependentModuleImpl struct {
	practice MainModule
}

func NewDependentInterface(practice MainModule) DependentModule {
	return &DependentModuleImpl{practice: practice}
}

func (d *DependentModuleImpl) DoPractice(printableString string) {
	fmt.Println(printableString)
	d.practice.RunMainModule(5)
}

